package logic

import (
	"context"
	"encoding/json"
	"github.com/hehanpeng/go-zero-resource/common/ctxdata"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/hehanpeng/go-zero-resource/message/internal/svc"
	"github.com/hehanpeng/go-zero-resource/message/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type KqSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKqSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KqSendLogic {
	return &KqSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KqSendLogic) KqSend(in *pb.KqSendReq) (*pb.Empty, error) {
	tracer := otel.GetTracerProvider().Tracer("kafka")
	spanCtx, span := tracer.Start(l.ctx, "send_msg_mq", trace.WithSpanKind(trace.SpanKindProducer))
	defer span.End()
	carrier := &propagation.HeaderCarrier{}
	otel.GetTextMapPropagator().Inject(spanCtx, carrier)
	msg := &ctxdata.MsgBody{
		Carrier: carrier,
		Msg:     in.Msg,
	}
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	err = l.svcCtx.KafkaTestPusher.Push(string(b))
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
