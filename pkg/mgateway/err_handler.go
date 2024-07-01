package mgateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zxcblog/rat-race/framework/gateway/render"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func StatusCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return 499
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

const (
	proxyFlag = "__succ__"
)

// ResponseHandler 返回消息处理
func ResponseHandler(ctx context.Context, res http.ResponseWriter, pbMsg proto.Message) error {

	return render.JSON{Data: Response{
		Code: http.StatusOK,
		Msg:  "",
		Data: pbMsg,
	}}.Render(res)
}

// ErrHandler 错误处理
func ErrHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, resWriter http.ResponseWriter, req *http.Request, err error) {
	convertErr := status.Convert(err)

	data := Response{
		Code: StatusCode(convertErr.Code()),
		Msg:  convertErr.Message(),
	}
	render.JSON{Data: data}.Render(resWriter)
}
