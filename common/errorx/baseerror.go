package errorx

import (
	"github.com/golang/protobuf/proto"
	"github.com/hehanpeng/go-zero-resource/sys/pb"
	"github.com/zeromicro/go-zero/core/mapping"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"net/http"
	"strconv"
)

const defaultCode = 400

const defaultErrorCode = -999

type CodeError struct {
	Code      uint32 `json:"code"`
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

type CodeErrorResponse struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

func NewCodeError(code uint32, errorCode int, msg string) error {
	return &CodeError{Code: code, ErrorCode: errorCode, Message: msg}
}

func NewEnumError(enum protoreflect.Enum) error {
	eCode, _ := proto.GetExtension(proto.MessageV1(enum.Descriptor().Values().ByNumber(enum.Number()).Options()), pb.E_Code)
	code, _ := strconv.ParseUint(mapping.Repr(eCode), 10, 32)
	eName, _ := proto.GetExtension(proto.MessageV1(enum.Descriptor().Values().ByNumber(enum.Number()).Options()), pb.E_Name)
	return &CodeError{Code: uint32(code), ErrorCode: int(enum.Number()), Message: mapping.Repr(eName)}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, defaultErrorCode, msg)
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
