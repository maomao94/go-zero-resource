package rpcserver

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hehanpeng/go-zero-resource/common/errorx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mapping"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*errorx.CodeError); ok {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
			metadata := make(map[string]string)
			metadata["errorCode"] = mapping.Repr(e.Code)
			metadata["message"] = e.Message
			errInfo := &errdetails.ErrorInfo{
				Reason:   e.Message,
				Domain:   "http://zero",
				Metadata: metadata,
			}
			var details []proto.Message
			details = append(details, errInfo)
			st, _ := status.New(codes.Internal, fmt.Sprintf("%d, %s", e.Code, e.Message)).WithDetails(details...)
			err = st.Err()
		} else {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
		}
	}
	return resp, err
}
