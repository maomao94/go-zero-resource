package kafka

import (
	"encoding/json"
	"github.com/hehanpeng/go-zero-resource/common/ctxdata"
	"github.com/hehanpeng/go-zero-resource/message/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	trace2 "github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
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
	var msg ctxdata.MsgBody
	if err := json.Unmarshal([]byte(value), &msg); err != nil {
		logx.Errorf(" consumer err : %v", err)
	} else {
		wireContext := otel.GetTextMapPropagator().Extract(ctx, msg.Carrier)
		tracer := otel.GetTracerProvider().Tracer(trace2.TraceName)
		_, span := tracer.Start(wireContext, "mq_consumer_msg", trace.WithSpanKind(trace.SpanKindProducer))
		defer span.End()
		logx.WithContext(wireContext).Infof("consumerOne Consumer, key: %+v msg:%+v", key, msg)
	}
	return nil
}
