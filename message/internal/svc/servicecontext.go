package svc

import (
	"github.com/hehanpeng/go-zero-resource/message/internal/config"
	"github.com/zeromicro/go-queue/kq"
)

type ServiceContext struct {
	Config    config.Config
	KafkaTest *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		KafkaTest: kq.NewPusher(c.Kafka.Brokers, c.Kafka.Topic),
	}
}
