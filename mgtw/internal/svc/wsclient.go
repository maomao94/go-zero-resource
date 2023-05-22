package svc

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hehanpeng/go-zero-resource/common/wsx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mapping"
	"github.com/zeromicro/go-zero/core/threading"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60
)

type Login struct {
	Seq    string
	Cmd    string
	AppId  uint32
	UserId string
	Client *Client
}

func (l *Login) GetKey() (key string) {
	key = GetUserKey(l.AppId, l.UserId)
	return
}

func GetUserKey(appId uint32, userId string) (key string) {
	key = fmt.Sprintf("%d_%s", appId, userId)
	return
}

type Client struct {
	SvcCtx        *ServiceContext
	Addr          string          // 客户端地址
	AppId         uint32          // 登录的平台Id app/web/ios
	UserId        string          // 用户Id，用户登录以后才有
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
	socket        *websocket.Conn // 用户连接
	send          chan []byte     // 待发送的数据
}

func NewClientCtx(svcCtx *ServiceContext, addr string, socket *websocket.Conn, firstTime uint64) (c *Client) {
	c = &Client{
		SvcCtx:        svcCtx,
		Addr:          addr,
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
		socket:        socket,
		send:          make(chan []byte, 100),
	}
	return
}

func (c *Client) closeSend() {
	close(c.send)
}

func (c *Client) GetKey() (key string) {
	key = GetUserKey(c.AppId, c.UserId)
	return
}

// 读取客户端数据
func (c *Client) Read() {
	defer func() {
		if r := recover(); r != nil {
			logx.Errorf("read error:%v", r)
		}
	}()
	defer func() {
		logx.Info("read send close")
		close(c.send)
	}()
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			logx.Errorf("socket error addr:%s:^err:%v", c.Addr, err)
			return
		}
		err = ProcessData(c, message)
		if err != nil {
			logx.Errorf("ProcessData error:%v", err)
			return
		}
	}
}

func ProcessData(c *Client, message []byte) (err error) {
	logx.Infof("ProcessData: len(message)=%d", len(message))
	ws := &wsx.WsRequest{}
	err = mapping.UnmarshalYamlBytes(message, ws)
	if err != nil {
		return
	}
	seq := ws.Seq
	cmd := ws.Cmd
	logx.Infof("%s-ProcessData: message:%s^size:%d", seq, ws.String(), len(message))

	switch cmd {
	case "login":
		loginReq := &wsx.LoginReq{}
		err = mapping.UnmarshalJsonMap(ws.Data, loginReq)
		if err != nil {
			return
		}
		login := &Login{
			Seq:    seq,
			Cmd:    cmd,
			AppId:  loginReq.AppId,
			UserId: loginReq.UserId,
			Client: c,
		}
		c.SvcCtx.ClientManager.PublishLogin(login)
	default:
		logx.Errorf("%s-cmd not found:%s", seq, cmd)
	}
	return
}

// 向客户端写数据
func (c *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			logx.Errorf("write error:%v", r)
		}
	}()
	defer func() {
		c.SvcCtx.ClientManager.PublishUnregister(c)
		c.socket.Close()
		logx.Info("write socket close")
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) SendSeqMsg(seq string, msg []byte) error {
	logx.Infof("%s-sendMsg msg:%s^size:%d", seq, msg, len(msg))
	var isSuccess bool
	if c == nil {
		return errors.New("client is nil")
	}
	threading.RunSafe(func() {
		c.send <- msg
		isSuccess = true
	})
	if !isSuccess {
		return errors.New("send msg fail")
	}
	return nil
}

func (c *Client) SendMsg(msg []byte) error {
	return c.SendSeqMsg("", msg)
}

func (c *Client) Heartbeat(currentTime uint64) {
	c.HeartbeatTime = currentTime
	return
}

func (c *Client) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.HeartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}
	return
}

func (c *Client) IsLogin() (isLogin bool) {
	if c.UserId != "" {
		isLogin = true
		return
	}
	return
}
