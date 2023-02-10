package svc

import (
	"github.com/honkkki/gomall/code/mall/service/user/api/internal/config"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
