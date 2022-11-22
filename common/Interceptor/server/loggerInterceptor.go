package server

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
	}
	return resp, err
}
