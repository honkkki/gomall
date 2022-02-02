package svc

import (
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/config"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/productclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
