package captcha

type CaptchaServer struct{}

func NewCaptchaServer() *CaptchaServer {
	return &CaptchaServer{}
}

// Get 获取验证码
//func (c *CaptchaServer) Get(ctx *gateway.Context) {
//resp := app.NewResponse(ctx)
//
//// 获取验证码
//id, b64, _, err := client.Captcha.Generate()
//if err != nil {
//	client.Log.ErrorF(ctx, "验证码获取错误：%s", err.Error())
//	resp.ToResponseError(app.ServerError.WithDetails("验证码获取错误"))
//	return
//}
//
//resp.ToResponse(gateway.H{"id": id, "captcha": b64})
//}
