package handler

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

// ErrHandler 错误处理
func ErrHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, resWriter http.ResponseWriter, req *http.Request, err error) {

	fmt.Println("现在是报错")

}
