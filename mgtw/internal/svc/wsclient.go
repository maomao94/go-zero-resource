package svc

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
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
	Socket        *websocket.Conn // 用户连接
	Send          chan []byte     // 待发送的数据
	AppId         uint32          // 登录的平台Id app/web/ios
	UserId        string          // 用户Id，用户登录以后才有
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
}

func NewClientCtx(svcCtx *ServiceContext, addr string, socket *websocket.Conn, firstTime uint64) (c *Client) {
	c = &Client{
		SvcCtx:        svcCtx,
		Addr:          addr,
		Socket:        socket,
		Send:          make(chan []byte, 100),
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
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
		close(c.Send)
	}()
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			logx.Errorf("socket 读取数组错误 addr: %s: %s", c.Addr, err)
			return
		}
		ProcessData(c, message)
	}
}

func ProcessData(c *Client, message []byte) {
	logx.Infof("收到数据: %s", string(message))
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
		c.Socket.Close()
		logx.Info("write socket close")
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) SendMsg(msg []byte) {
	if c == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			logx.Errorf("SendMsg: %s")
		}
	}()
	c.Send <- msg
}

func (c *Client) close() {
	close(c.Send)
}

func (c *Client) Login(appId uint32, userId string, loginTime uint64) {
	c.AppId = appId
	c.UserId = userId
	c.LoginTime = loginTime
	// 登录成功=心跳一次
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
	// 用户登录了
	if c.UserId != "" {
		isLogin = true
		return
	}
	return
}
