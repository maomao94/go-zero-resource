package interceptor

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/hehanpeng/go-zero-resource/common/errorx"
	"github.com/hehanpeng/go-zero-resource/hello/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*errorx.CodeError); ok {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
			var details []proto.Message
			detail := &pb.ErrorDetail{
				ErrorCode: int32(e.ErrorCode),
				Message:   e.Error(),
			}
			details = append(details, detail)
			st, _ := status.New(codes.Code(e.Code), e.Error()).WithDetails(details...)
			err = st.Err()
		} else {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
		}

	}
	return resp, err
}
