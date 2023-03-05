package svc

import (
	"github.com/honkkki/gomall/code/mall/service/order/rpc/orderclient"
	"github.com/honkkki/gomall/code/mall/service/pay/model"
	"github.com/honkkki/gomall/code/mall/service/pay/rpc/internal/config"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	PayModel model.PayModel

	UserRpcClient  userclient.User
	OrderRpcClient orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		PayModel:       model.NewPayModel(conn, c.CacheRedis),
		UserRpcClient:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpcClient: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
