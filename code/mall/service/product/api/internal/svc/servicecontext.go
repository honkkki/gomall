package svc

import (
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/config"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/productclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	ProductRpcClient productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		ProductRpcClient: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
