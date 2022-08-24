package gtw

import (
	"context"
	"gtw/resource/resource"

	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"

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
	pingResp, err := l.svcCtx.ResourceRpc.Ping(l.ctx, &resource.Empty{})
	if err != nil {
		return nil, err
	}
	return &types.PingReply{
		Msg: pingResp.Msg,
	}, nil
}
