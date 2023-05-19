package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/net/context"
	"runtime/debug"

	"github.com/gorilla/websocket"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60
)

// 用户登录
type login struct {
	AppId  uint32
	UserId string
	Client *Client
}

// GetKey 获取 key
func (l *login) GetKey() (key string) {
	key = GetUserKey(l.AppId, l.UserId)
	return
}

// 获取用户key
func GetUserKey(appId uint32, userId string) (key string) {
	key = fmt.Sprintf("%d_%s", appId, userId)
	return
}

// 用户连接
type Client struct {
	Ctx           context.Context // 上下文
	Addr          string          // 客户端地址
	Socket        *websocket.Conn // 用户连接
	Send          chan []byte     // 待发送的数据
	AppId         uint32          // 登录的平台Id app/web/ios
	UserId        string          // 用户Id，用户登录以后才有
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
}

// 初始化
func NewClientCtx(ctx context.Context, addr string, socket *websocket.Conn, firstTime uint64) (client *Client) {
	client = &Client{
		Ctx:           ctx,
		Addr:          addr,
		Socket:        socket,
		Send:          make(chan []byte, 100),
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
	}
	return
}

// GetKey 获取 key
func (c *Client) GetKey() (key string) {
	//key = GetUserKey(c.AppId, c.UserId)
	return
}

// 读取客户端数据
func (c *Client) Read(ctx context.Context, svcCtx *ServiceContext) {
	defer func() {
		if r := recover(); r != nil {
			logx.WithContext(ctx).Info("write stop", string(debug.Stack()), r)
		}
	}()
	defer func() {
		logx.WithContext(ctx).Info("读取客户端数据 关闭send", c)
		close(c.Send)
	}()
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			logx.WithContext(ctx).Error("读取客户端数据 错误", c.Addr, err)
			return
		}
		// 处理程序
		logx.WithContext(ctx).Info("读取客户端数据 处理:", string(message))
		//ProcessData(c, message)
		svcCtx.ClientManager.Login <- &login{
			AppId:  111,
			UserId: "2222",
			Client: c,
		}
	}
}

// 向客户端写数据
func (c *Client) Write(ctx context.Context, svcCtx *ServiceContext) {
	defer func() {
		if r := recover(); r != nil {
			logx.WithContext(ctx).Info("write stop", string(debug.Stack()), r)
		}
	}()
	defer func() {
		svcCtx.ClientManager.Unregister <- c
		c.Socket.Close()
		logx.WithContext(ctx).Info("Client发送数据 defer", c)
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// 发送数据错误 关闭连接
				logx.WithContext(ctx).Info("Client发送数据 关闭连接", c.Addr, "ok", ok)
				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

// SendMsg 发送数据
func (c *Client) SendMsg(msg []byte) {
	if c == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()
	c.Send <- msg
}

// close 关闭客户端连接
func (c *Client) close() {
	close(c.Send)
}

// 用户登录
func (c *Client) Login(appId uint32, userId string, loginTime uint64) {
	c.AppId = appId
	c.UserId = userId
	c.LoginTime = loginTime
	// 登录成功=心跳一次
	c.Heartbeat(loginTime)
}

// 用户心跳
func (c *Client) Heartbeat(currentTime uint64) {
	c.HeartbeatTime = currentTime
	return
}

// 心跳超时
func (c *Client) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.HeartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}
	return
}

// 是否登录了
func (c *Client) IsLogin() (isLogin bool) {
	// 用户登录了
	if c.UserId != "" {
		isLogin = true
		return
	}
	return
}
