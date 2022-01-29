package svc

import (
	"github.com/honkkki/gomall/code/mall/service/user/api/internal/config"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/userclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
