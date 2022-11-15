package kafka

import (
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/net/context"
	"gtw/message/internal/svc"
)

type KafkaTest struct {
	svcCtx *svc.ServiceContext
}

func NewKafkaTest(svcCtx *svc.ServiceContext) *KafkaTest {
	return &KafkaTest{
		svcCtx: svcCtx,
	}
}

func (l KafkaTest) Consume(key, value string) error {
	ctx := context.Background()
	logx.WithContext(ctx).Infow("KafkaTest", logx.Field("key", key))
	logx.WithContext(ctx).Infow("KafkaTest", logx.Field("value", value))
	return nil
}
