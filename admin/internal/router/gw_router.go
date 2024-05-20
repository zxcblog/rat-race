package router

import (
	"context"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zxcblog/rat-race/internal/client"
	"github.com/zxcblog/rat-race/internal/router/pb/user"
	"github.com/zxcblog/rat-race/internal/server/captcha"
	"github.com/zxcblog/rat-race/pkg/gateway"
	"github.com/zxcblog/rat-race/pkg/swagger"
	grpc2 "google.golang.org/grpc"
	"net/http"
	"path"
	"strings"
)

func GWRouter() (*gateway.GWBuild, error) {

	opts := []gateway.Options{gateway.WithConfig(client.GwConf)}

	// dev模式运行时， 加载swagger
	if client.GwConf.RunMode == gateway.DevMod {
		// https://www.cnblogs.com/zhangmingcheng/p/16349561.html
		// go-bindata --nocompress -pkg swagger -o pkg/swagger/swagger.go pkg/swagger/ui/...
		mux := http.NewServeMux()

		// 加载swagger.json文件
		mux.HandleFunc("/swagger-file/", func(writer http.ResponseWriter, request *http.Request) {
			name := path.Join("internal/router/swagger", strings.TrimPrefix(request.URL.Path, "/swagger-file/"))
			http.ServeFile(writer, request, name)
		})

		fileServer := http.FileServer(&assetfs.AssetFS{
			Asset:    swagger.Asset,
			AssetDir: swagger.AssetDir,
			Prefix:   "pkg/swagger/ui",
		})
		prefix := "/swagger/"
		mux.Handle(prefix, http.StripPrefix(prefix, fileServer))

		opts = append(opts, gateway.WithSwagger("/swagger", mux))
	}

	server := gateway.NewGWBuild(opts...)
	server.RegisterServer(
		func(ctx context.Context, mux *runtime.ServeMux, conn *grpc2.ClientConn) error {
			if err := user.RegisterUserHandler(ctx, mux, conn); err != nil {
				return err
			}
			return nil
		},
	)

	// 验证码操作
	catpchaServer := captcha.NewCaptchaServer()
	server.GET("/v1/captcha", catpchaServer.Get)

	return server, nil
}
