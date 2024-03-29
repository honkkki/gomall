package svc

import (
	"github.com/honkkki/gomall/code/mall/service/product/model"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
	RedisClient  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn, c.RedisCache),
		RedisClient:  redis.New(c.RedisConf.Addr),
	}
}
