package errorx

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/mapping"
	"github.com/zeromicro/go-zero/core/trace"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"net/http"
	"strconv"
)

const defaultErrorCode = 999

type CodeError struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

type CodeErrorResponse struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	traceId   string `json:"traceId"`
}

func New(errCode int, message string, traceId string) *CodeErrorResponse {
	return &CodeErrorResponse{
		ErrorCode: errCode,
		Message:   message,
	}
}

func Default() *CodeErrorResponse {
	return &CodeErrorResponse{
		ErrorCode: defaultErrorCode,
		Message:   "未知错误",
	}
}

func NewCodeError(errorCode int, msg string) error {
	return &CodeError{ErrorCode: errorCode, Message: msg}
}

func NewEnumError(enum protoreflect.Enum) error {
	eName, _ := proto.GetExtension(proto.MessageV1(enum.Descriptor().Values().ByNumber(enum.Number()).Options()), E_Name)
	name := fmt.Sprintf("%s", mapping.Repr(eName))
	return &CodeError{ErrorCode: int(enum.Number()), Message: name}
}

func NewEnumErrorf(enum protoreflect.Enum, wrap string) error {
	err := NewEnumError(enum)
	if true {
		if e, ok := err.(*CodeError); ok {
			e.Message = fmt.Sprintf("%s, %s", err.Error(), wrap)
		}
	}
	return err
}

func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		ErrorCode: e.ErrorCode,
		Message:   e.Message,
	}
}

func CodeFromGrpcError(err error) int {
	code := status.Code(err)
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument, codes.FailedPrecondition, codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.AlreadyExists, codes.Aborted:
		return http.StatusConflict
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.Internal, codes.DataLoss, codes.Unknown:
		return http.StatusInternalServerError
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	}

	return http.StatusInternalServerError
}

func IsGrpcError(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(interface {
		GRPCStatus() *status.Status
	})

	return ok
}

func FromError(ctx context.Context, err error) *CodeErrorResponse {
	traceID := trace.TraceIDFromContext(ctx)
	if err == nil {
		return &CodeErrorResponse{
			ErrorCode: defaultErrorCode,
			Message:   "err is nil",
			traceId:   traceID,
		}
	}
	gs, ok := status.FromError(err)
	if ok {
		for _, detail := range gs.Details() {
			switch d := detail.(type) {
			case *errdetails.ErrorInfo:
				metadata := d.Metadata
				ec, _ := metadata["errorCode"]
				errorCode, e := strconv.ParseInt(ec, 10, 32)
				if e != nil {
					return Default()
				}
				message, _ := metadata["message"]
				return New(int(errorCode), message, traceID)
			}
		}
	}
	return &CodeErrorResponse{
		ErrorCode: defaultErrorCode,
		Message:   err.Error(),
		traceId:   traceID,
	}
}
