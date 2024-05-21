package gateway

import (
	"testing"
)

func TestNew(t *testing.T) {

	engine := New()

	group := engine.Group("/v1")
	group.GET("/hello", func(ctx *Context) {
		ctx.JSON(404, "hello word")
	})

	group.StaticFile("/config", "./config/config.yaml")
	group.StaticFileFS("/config/{path}", "path", Dir("./config", true))

	engine.Run()
}
