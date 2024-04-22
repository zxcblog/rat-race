package grpc

type Options interface {
	apply(build *GRPCBuild)
}

type OptionFunc func(*GRPCBuild)

func (f OptionFunc) apply(b *GRPCBuild) {
	f(b)
}
