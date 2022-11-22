package gtw

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/hello/pb"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingHelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingHelloLogic {
	return &PingHelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingHelloLogic) PingHello() (resp *types.PingReply, err error) {
	pingResp, err := l.svcCtx.HelloRpc.Ping(l.ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	return &types.PingReply{Msg: pingResp.Msg}, nil
}
