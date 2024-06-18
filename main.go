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
	router.GatewayRouter()

	client.Shutdown.Close()
}
