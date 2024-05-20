package gateway

import "net/http"

type Options interface {
	apply(build *GWBuild)
}

type OptionFunc func(*GWBuild)

func (f OptionFunc) apply(b *GWBuild) {
	f(b)
}

func WithConfig(conf *Config) Options {
	return OptionFunc(func(build *GWBuild) {
		build.conf = conf
	})
}

func WithSwagger(swaggerPrefix string, handler http.Handler) Options {
	return OptionFunc(func(gw *GWBuild) {
		gw.swagger = true
		gw.swaggerPrefix = swaggerPrefix
		gw.swaggerHandler = handler
	})
}
