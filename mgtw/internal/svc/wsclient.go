package svc

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60
)

type login struct {
	AppId  uint32
	UserId string
	Client *Client
}

func (l *login) GetKey() (key string) {
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

func (c *Client) GetKey() (key string) {
	key = GetUserKey(c.AppId, c.UserId)
	return
}

// 读取客户端数据
func (c *Client) Read() {
	defer func() {
		if r := recover(); r != nil {
			logx.Errorf("read error: %s", r)
		}
	}()
	defer func() {
		logx.Info("read send close")
		close(c.send)
	}()
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			logx.Errorf("socket 读取数组错误 addr: %s: %s", c.Addr, err)
			return
		}
		ProcessData(c, message)
	}
}

func ProcessData(c *Client, message []byte) {
	logx.Infof("ProcessData: %s", string(message))
	c.SvcCtx.ClientManager.Login <- &login{
		AppId:  111,
		UserId: string(message),
		Client: c,
	}
}

// 向客户端写数据
func (c *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			logx.Errorf("write error: %s", r)
		}
	}()
	defer func() {
		c.SvcCtx.ClientManager.Unregister <- c
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

func (c *Client) SendMsg(msg []byte) error {
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

func (c *Client) closeSend() {
	close(c.send)
}

func (c *Client) Login(appId uint32, userId string, loginTime uint64) {
	c.AppId = appId
	c.UserId = userId
	c.LoginTime = loginTime
	c.Heartbeat(loginTime)
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
