package svc

import (
	"github.com/honkkki/gomall/code/mall/service/order/api/internal/config"
	"github.com/honkkki/gomall/code/mall/service/order/rpc/orderclient"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/productclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	OrderRpcClient   orderclient.Order
	ProductRpcClient productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		OrderRpcClient:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpcClient: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
