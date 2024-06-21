package captcha

import (
	"github.com/zxcblog/rat-race/app/client"
	"github.com/zxcblog/rat-race/pkg/mgateway"
	"net/http"
)

type CaptchaServer struct{}

func NewCaptchaServer() *CaptchaServer {
	return &CaptchaServer{}
}

// Get 获取验证码
func (c *CaptchaServer) Get(ctx *mgateway.Context) {
	//resp := app.NewResponse(ctx)

	// 获取验证码
	id, b64, _, err := client.Captcha.Generate()
	if err != nil {
		client.Logger.ErrorFWithCtx(ctx, "验证码获取错误：%s", err.Error())
		//resp.ToResponseError(app.ServerError.WithDetails("验证码获取错误"))
		return
	}

	ctx.JSON(http.StatusOK, mgateway.H{"code": 20000, "id": id, "captcha": b64})
}
