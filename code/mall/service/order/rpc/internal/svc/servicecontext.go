package svc

import (
	"github.com/honkkki/gomall/code/mall/service/order/model"
	"github.com/honkkki/gomall/code/mall/service/order/rpc/internal/config"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/productclient"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/types/product"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/types/user"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel

	UserRpcClient    user.UserClient
	ProductRpcClient product.ProductClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:           c,
		OrderModel:       model.NewOrderModel(conn, c.RedisCache),
		UserRpcClient:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpcClient: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
