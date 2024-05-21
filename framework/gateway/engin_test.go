package gateway

import (
	"testing"
)

func TestNew(t *testing.T) {

	engine := New()

	engine.GET("/hello", func(ctx *Context) {
		ctx.JSON(404, "hello word")
	})

	engine.StaticFile("/config", "./config/config.yaml")
	engine.StaticFileFS("/config/{path}", "path", Dir("./config", true))

	engine.Run()
}
