package svc

import (
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
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
	threading.GoSafe(func() {
		logx.Info("start webSocket listener")
		m.StartListener()
	})
	return
}

// 管道处理程序
func (manager *ClientManager) StartListener() {
	for {
		select {
		case client := <-manager.Register:
			// 建立连接事件
			manager.EventRegister(client)
		case login := <-manager.Login:
			// 用户登录
			manager.EventLogin(login)
		case client := <-manager.Unregister:
			// 断开连接事件
			manager.EventUnregister(client)
		case message := <-manager.Broadcast:
			// 广播事件
			clients := manager.GetClients()
			for conn := range clients {
				err := conn.SendMsg(message)
				if err != nil {
					logx.Errorf("广播消息失败:%v", err)
					continue
				}
			}
		}
	}
}

func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)
	logx.Infof("eventRegister addr:%s", client.Addr)
}

func (manager *ClientManager) PublishRegister(client *Client) {
	threading.RunSafe(func() {
		manager.Register <- client
	})
}

func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	manager.Clients[client] = true
}

func (manager *ClientManager) EventLogin(l *login) {
	client := l.Client
	// 连接存在，在添加
	if manager.InClient(client) {
		userKey := l.GetKey()
		manager.AddUsers(userKey, l.Client)
	}
	logx.Infof("eventLogin addr:%s^appId:%d^userId:%s", client.Addr, l.AppId, l.UserId)
	manager.Broadcast <- []byte("有用户登录了")
}

func (manager *ClientManager) EventUnregister(client *Client) {

}

func (manager *ClientManager) GetClientsLen() (clientsLen int) {
	clientsLen = len(manager.Clients)
	return
}

func (manager *ClientManager) GetUsersLen() (userLen int) {
	userLen = len(manager.Users)
	return
}

func (manager *ClientManager) InClient(client *Client) (ok bool) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	// 连接存在，在添加
	_, ok = manager.Clients[client]
	return
}

func (manager *ClientManager) AddUsers(key string, client *Client) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()
	manager.Users[key] = client
}

func (manager *ClientManager) GetClients() (clients map[*Client]bool) {
	clients = make(map[*Client]bool)
	manager.ClientsRange(func(client *Client, value bool) (result bool) {
		clients[client] = value
		return true
	})
	return
}

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

func (manager *ClientManager) DelClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	if _, ok := manager.Clients[client]; ok {
		delete(manager.Clients, client)
	}
}

func (manager *ClientManager) GetUserClient(appId uint32, userId string) (client *Client) {
	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()
	userKey := GetUserKey(appId, userId)
	if value, ok := manager.Users[userKey]; ok {
		client = value
	}
	return
}

func (manager *ClientManager) DelUsers(client *Client) (result bool) {
	manager.UserLock.Lock()
	defer manager.UserLock.Unlock()
	key := GetUserKey(client.AppId, client.UserId)
	if value, ok := manager.Users[key]; ok {
		// 判断是否为相同的用户
		if value.Addr != client.Addr {
			return
		}
		delete(manager.Users, key)
		result = true
	}
	return
}

func (manager *ClientManager) GetUserKeys() (userKeys []string) {
	userKeys = make([]string, 0)
	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()
	for key := range manager.Users {
		userKeys = append(userKeys, key)
	}
	return
}

func (manager *ClientManager) GetUserList(appId uint32) (userList []string) {
	userList = make([]string, 0)
	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()
	for _, v := range manager.Users {
		if v.AppId == appId {
			userList = append(userList, v.UserId)
		}
	}
	return
}

func (manager *ClientManager) GetUserClients() (clients []*Client) {
	clients = make([]*Client, 0)
	manager.UserLock.RLock()
	defer manager.UserLock.RUnlock()
	for _, v := range manager.Users {
		clients = append(clients, v)
	}
	return
}

func (manager *ClientManager) sendAll(message []byte, ignoreClient *Client) {
	clients := manager.GetUserClients()
	for _, conn := range clients {
		if conn != ignoreClient {
			err := conn.SendMsg(message)
			if err != nil {
				logx.Errorf("sendAll key:%s^err:%v", conn.GetKey(), err)
			}
		}
	}
}

func (manager *ClientManager) sendAppIdAll(message []byte, appId uint32, ignoreClient *Client) {
	clients := manager.GetUserClients()
	for _, conn := range clients {
		if conn != ignoreClient && conn.AppId == appId {
			err := conn.SendMsg(message)
			if err != nil {
				logx.Errorf("sendAppIdAll key:%s^err:%v", conn.GetKey(), err)
			}
		}
	}
}
