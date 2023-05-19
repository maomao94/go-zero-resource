package svc

import (
	"github.com/hehanpeng/go-zero-resource/common/Interceptor/rpcclient"
	"github.com/hehanpeng/go-zero-resource/message/internal/config"
	"github.com/hehanpeng/go-zero-resource/mgtw/mgtw"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"math/rand"
	"sync"
)

const (
	subsetSize = 32
)

var Conn4UniqueList []*zrpc.Client
var Conn4UniqueListMtx sync.RWMutex
var IsUpdateStart bool
var IsUpdateStartMtx sync.RWMutex

type ServiceContext struct {
	Config          config.Config
	KafkaTestPusher *kq.Pusher
	PubContainer    *PubContainer
	MGtwRpc         mgtw.Mgtw
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		KafkaTestPusher: kq.NewPusher(c.Kafka.Brokers, c.Kafka.Topic),
		PubContainer:    NewEtcdPubContainer(c.MGtwRpcConf),
		MGtwRpc: mgtw.NewMgtw(zrpc.MustNewClient(
			c.MGtwRpcConf, zrpc.WithUnaryClientInterceptor(rpcclient.UnaryMetadataInterceptor))),
	}
}

type PubContainer struct {
	PubMap map[string]mgtw.Mgtw
	lock   sync.Mutex
}

func NewEtcdPubContainer(c zrpc.RpcClientConf) *PubContainer {
	p := &PubContainer{
		PubMap: make(map[string]mgtw.Mgtw),
	}
	if len(c.Endpoints) != 0 {
		err := p.getConn4Direct(c)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	err := p.getConn4Etcd(c)
	if err != nil {
		log.Fatal(err.Error())
	}
	return p
}

func (p *PubContainer) getConn4Etcd(c zrpc.RpcClientConf) error {
	sub, err := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
	if err != nil {
		return err
	}
	update := func() {
		var add []string
		var remove []string
		p.lock.Lock()
		m := make(map[string]any)
		for _, val := range subset(sub.Values(), subsetSize) {
			m[val] = true
		}
		for k, _ := range p.PubMap {
			if _, ok := m[k]; !ok {
				remove = append(remove, k)
			}
		}
		for k, _ := range m {
			if _, ok := p.PubMap[k]; !ok {
				add = append(add, k)
			}
		}
		for _, val := range add {
			endpoints := make([]string, 1)
			endpoints[0] = val
			c.Endpoints = endpoints
			pub := mgtw.NewMgtw(zrpc.MustNewClient(
				c, zrpc.WithUnaryClientInterceptor(rpcclient.UnaryMetadataInterceptor)))
			p.PubMap[val] = pub
		}
		for _, val := range remove {
			delete(p.PubMap, val)
		}
		p.lock.Unlock()
	}
	sub.AddListener(update)
	update()
	return nil
}

func (p *PubContainer) getConn4Direct(c zrpc.RpcClientConf) error {
	p.lock.Lock()
	for _, val := range c.Endpoints {
		if _, ok := p.PubMap[val]; ok {
			continue
		}
		endpoints := make([]string, 1)
		endpoints[0] = val
		c.Endpoints = endpoints
		pub := mgtw.NewMgtw(zrpc.MustNewClient(
			c, zrpc.WithUnaryClientInterceptor(rpcclient.UnaryMetadataInterceptor)))
		p.PubMap[val] = pub
	}
	p.lock.Unlock()
	return nil
}

func subset(set []string, sub int) []string {
	rand.Shuffle(len(set), func(i, j int) {
		set[i], set[j] = set[j], set[i]
	})
	if len(set) <= sub {
		return set
	}

	return set[:sub]
}
