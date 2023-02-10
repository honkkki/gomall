package svc

import (
	"github.com/honkkki/gomall/code/mall/service/order/api/internal/config"
	"github.com/honkkki/gomall/code/mall/service/order/rpc/orderclient"
	"github.com/honkkki/gomall/code/mall/service/order/rpc/types/order"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	OrderRpcClient order.OrderClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		OrderRpcClient: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
