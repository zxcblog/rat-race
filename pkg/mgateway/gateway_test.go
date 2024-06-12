package mgateway

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {

	engine := New()

	group := engine.Group("/v1")

	group.Use(Cors)

	group.GET("/hello", func(ctx *Context) {
		ctx.JSON(200, "hello word")
	})

	group.StaticFile("/config", "./config/config.yaml")
	group.StaticFileFS("/config/{path}", "path", Dir("./config", true))

	engine.Run()
}

func Cors(c *Context) {
	method := c.Request.Method
	origin := c.Request.Header.Get("Origin")
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
	}
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}
