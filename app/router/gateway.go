package router

func GatewayRouter() {
	//ctx := context.Background()
	//route := mgateway.New(client.Config.Server.GatewayConf, client.Logger)
	//
	//route.Use(handler.Cors)
	//// 注册自定义请求地址
	//captS := captcha.NewCaptchaServer()
	//route.GET("/v1/captcha/img", captS.Get)
	//
	//// 注册grpc http服务
	//grpcConn := mgrpc.GetConnByEtcd(client.Config.Server.Name, client.Etcd, client.Logger)
	//if err := user.RegisterUserHandler(ctx, route.GetServerMux(), grpcConn); err != nil {
	//	panic(err.Error())
	//}
	//
	//route.Run()
	//// 注册关闭服务
	//client.Shutdown.Register(func() {
	//	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	//	defer cancel()
	//	if err := route.Close(ctx); err != nil {
	//		client.Logger.ErrorF("gateway服务关闭失败:", err)
	//	}
	//
	//	client.Logger.InfoF("gateway服务退出")
	//})
}
