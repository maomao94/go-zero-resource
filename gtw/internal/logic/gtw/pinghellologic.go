package gtw

import (
	"context"

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
	return
}