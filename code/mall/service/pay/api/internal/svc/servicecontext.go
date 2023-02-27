package svc

import (
	"github.com/honkkki/gomall/code/mall/service/pay/api/internal/config"
	"github.com/honkkki/gomall/code/mall/service/pay/rpc/payclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PayRpcClient payclient.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		PayRpcClient: payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
