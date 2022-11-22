package interceptor

import (
	"context"
	"github.com/hehanpeng/go-zero-resource/common/ctxdata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
)

func UnaryMetadataInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	md.Set(ctxdata.CtxKeyUserId, strconv.FormatInt(ctxdata.GetUserIdFromCtx(ctx, false), 10))
	metaCtx := metadata.NewOutgoingContext(ctx, md)
	return invoker(metaCtx, method, req, reply, cc, opts...)
}

func StreamTracingInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
	method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	md.Set(ctxdata.CtxKeyUserId, strconv.FormatInt(ctxdata.GetUserIdFromCtx(ctx, false), 10))
	metaCtx := metadata.NewOutgoingContext(ctx, md)
	return streamer(metaCtx, desc, cc, method, opts...)
}
