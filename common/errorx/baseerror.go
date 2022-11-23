package errorx

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
}

func New(errCode int, message string) *CodeErrorResponse {
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

//func NewEnumError(enum protoreflect.Enum) error {
//	eCode, _ := proto.GetExtension(proto.MessageV1(enum.Descriptor().Values().ByNumber(enum.Number()).Options()), pb.E_Code)
//	code, _ := strconv.ParseUint(mapping.Repr(eCode), 10, 32)
//	eName, _ := proto.GetExtension(proto.MessageV1(enum.Descriptor().Values().ByNumber(enum.Number()).Options()), pb.E_Name)
//	return &CodeError{Code: uint32(code), ErrorCode: int(enum.Number()), Message: mapping.Repr(eName)}
//}

//func NewEnumErrorf(enum protoreflect.Enum, wrap string) error {
//	eBool, _ := proto.GetExtension(proto.MessageV1(enum.Descriptor().Values().ByNumber(enum.Number()).Options()), pb.E_Bool)
//	bool, _ := strconv.ParseBool(mapping.Repr(eBool))
//	err := NewEnumError(enum)
//	if bool {
//		append := fmt.Sprintf("%s^", err.Error())
//		if e, ok := err.(*CodeError); ok {
//			e.Message = fmt.Sprintf("%s%s", append, wrap)
//		}
//	}
//	return err
//}

//func NewDefaultError(msg string) error {
//	return NewCodeError(defaultCode, defaultErrorCode, msg)
//}

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

func FromError(err error) *CodeErrorResponse {
	if err == nil {
		return &CodeErrorResponse{
			ErrorCode: defaultErrorCode,
			Message:   "err is nil",
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
				return New(int(errorCode), message)
			}
		}
	}
	return &CodeErrorResponse{
		ErrorCode: defaultErrorCode,
		Message:   err.Error(),
	}
}
