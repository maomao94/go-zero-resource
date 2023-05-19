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
	// DirectScheme stands for direct scheme.
	DirectScheme = "direct"
	// DiscovScheme stands for discov scheme.
	DiscovScheme = "discov"
	// EtcdScheme stands for etcd scheme.
	EtcdScheme = "etcd"
	// KubernetesScheme stands for k8s scheme.
	KubernetesScheme = "k8s"
	// EndpointSepChar is the separator cha in endpoints.
	EndpointSepChar = ','

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
	MGtwRpcList []mgtw.Mgtw
}

func NewEtcdPubContainer(c zrpc.RpcClientConf) *PubContainer {
	p := &PubContainer{}
	if len(c.Endpoints) != 0 {
		err := p.getConn4UniqueCfg(c)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	err := p.getConn4UniqueEtcd(c)
	if err != nil {
		log.Fatal(err.Error())
	}
	return p
}

func (p *PubContainer) getConn4UniqueEtcd(c zrpc.RpcClientConf) error {
	sub, err := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
	if err != nil {
		return err
	}
	update := func() {
		updatePub := make([]mgtw.Mgtw, 0)
		for _, val := range subset(sub.Values(), subsetSize) {
			endpoints := make([]string, 1)
			endpoints[0] = val
			c.Endpoints = endpoints
			mGtwRpc := mgtw.NewMgtw(zrpc.MustNewClient(
				c, zrpc.WithUnaryClientInterceptor(rpcclient.UnaryMetadataInterceptor)))
			updatePub = append(updatePub, mGtwRpc)
		}
		p.MGtwRpcList = updatePub
	}
	sub.AddListener(update)
	update()
	return nil
}

func (p *PubContainer) getConn4UniqueCfg(c zrpc.RpcClientConf) error {
	pub := make([]mgtw.Mgtw, 0)
	for _, val := range c.Endpoints {
		endpoints := make([]string, 1)
		endpoints[0] = val
		c.Endpoints = endpoints
		mGtwRpc := mgtw.NewMgtw(zrpc.MustNewClient(
			c, zrpc.WithUnaryClientInterceptor(rpcclient.UnaryMetadataInterceptor)))
		pub = append(pub, mGtwRpc)
	}
	p.MGtwRpcList = pub
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
