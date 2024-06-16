package main

import (
	"github.com/zxcblog/rat-race/app/client"
	"github.com/zxcblog/rat-race/app/router"
)

func main() {
	if err := client.Init("./config/config.yaml"); err != nil {
		panic(err.Error())
	}

	router.GrpcRouter()

	//i := 0
	//for i < 5 {
	//	client.Logger.DebugF("准备发送请求")
	//	time.Sleep(5 * time.Second)
	//	i++
	//
	//	resp, err := client.RatRaceMicro.Login(context.Background(), &user.LoginReq{
	//		Type:     0,
	//		Account:  "admin",
	//		Password: "adminadmin",
	//		Captcha:  "zxcf",
	//	})
	//
	//	if err != nil {
	//		client.Logger.ErrorF(err.Error())
	//		return
	//	}
	//	client.Logger.InfoF("打印请求信息%+v", resp)
	//}

	client.Shutdown.Close()
}
