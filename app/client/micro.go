package client

import (
	"github.com/zxcblog/rat-race/app/pb/user"
	"github.com/zxcblog/rat-race/pkg/mgrpc"
)

type ratRaceMicro struct {
	user.UserClient
}

func newRatRaceMicro() ratRaceMicro {
	conn := mgrpc.GetConnByEtcd(Config.Server.Name, Etcd, Logger)

	return ratRaceMicro{
		UserClient: user.NewUserClient(conn),
	}
}
