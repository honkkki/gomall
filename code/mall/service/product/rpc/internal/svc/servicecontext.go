package svc

import (
	"github.com/honkkki/gomall/code/mall/service/product/model"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/internal/config"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
	DBEngine     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)

	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn, c.CacheRedis),
		DBEngine:     db,
	}
}
