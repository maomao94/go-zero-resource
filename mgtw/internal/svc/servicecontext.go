package svc

import (
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type ServiceContext struct {
	Config        config.Config
	ClientManager *ClientManager
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ClientManager: NewClientManager(),
	}
}

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
		case login := <-manager.Login:
			// 用户登录
			manager.EventLogin(login)
		case conn := <-manager.Unregister:
			// 断开连接事件
			manager.EventUnregister(conn)
		case message := <-manager.Broadcast:
			// 广播事件
			clients := manager.GetClients()
			for conn := range clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
				}
			}
		}
	}
}

// 用户建立连接事件
func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)
	logx.WithContext(client.Ctx).Info("EventRegister 用户建立连接", client.Addr)
}

// 添加客户端
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	manager.Clients[client] = true
}

func (manager *ClientManager) EventLogin(l *login) {

}

func (manager *ClientManager) GetClients() (clients map[*Client]bool) {
	clients = make(map[*Client]bool)
	manager.ClientsRange(func(client *Client, value bool) (result bool) {
		clients[client] = value
		return true
	})
	return
}

// 遍历
func (manager *ClientManager) ClientsRange(f func(client *Client, value bool) (result bool)) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	for key, value := range manager.Clients {
		result := f(key, value)
		if result == false {
			return
		}
	}
	return
}

func (manager *ClientManager) EventUnregister(conn *Client) {

}
