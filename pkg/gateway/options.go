package gateway

type Options interface {
	apply(build *GWBuild)
}

type OptionFunc func(*GWBuild)

func (f OptionFunc) apply(b *GWBuild) {
	f(b)
}
