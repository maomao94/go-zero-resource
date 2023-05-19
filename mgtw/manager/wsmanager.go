package manager

import (
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type ClientManager struct {
	Clients     map[*Client]bool   // 全部的连接
	ClientsLock sync.RWMutex       // 读写锁
	Users       map[string]*Client // 登录的用户 // appId+uuid
	UserLock    sync.RWMutex       // 读写锁
	Register    chan *Client       // 连接连接处理
	Login       chan *login        // 用户登录处理
	Unregister  chan *Client       // 断开连接处理程序
	Broadcast   chan []byte        // 广播 向全部成员发送数据
}

func NewClientManager() (m *ClientManager) {
	m = &ClientManager{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
		Login:      make(chan *login, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}
	go m.start()
	return
}

// 管道处理程序
func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			// 建立连接事件
			manager.EventRegister(conn)
			//case login := <-manager.Login:
			//	// 用户登录
			//	//manager.EventLogin(login)
			//case conn := <-manager.Unregister:
			//	// 断开连接事件
			//	//manager.EventUnregister(conn)
			//case message := <-manager.Broadcast:
			//	// 广播事件
			//	clients := manager.GetClients()
			//	for conn := range clients {
			//		select {
			//		case conn.Send <- message:
			//		default:
			//			close(conn.Send)
			//		}
			//	}
		}
	}
}

// 用户建立连接事件
func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)
	logx.Info("EventRegister 用户建立连接", client.Addr)
}

// 添加客户端
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	manager.Clients[client] = true
}
