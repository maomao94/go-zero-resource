package gtw

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/resource/pb"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.PingReply, err error) {
	pingResp, err := l.svcCtx.ResourceRpc.Ping(l.ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	return &types.PingReply{
		Msg: pingResp.Msg,
	}, nil
}
