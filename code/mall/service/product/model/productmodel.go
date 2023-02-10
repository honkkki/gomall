package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		UpdateStock(ctx context.Context, pid int64, num int32) error
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c),
	}
}

func (c *customProductModel) UpdateStock(ctx context.Context, pid int64, num int32) error {
	productIdKey := fmt.Sprintf("%s%v", cacheProductIdPrefix, pid)
	_, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set stock = stock - ? where `id` = ? and `stock` > 0", c.table)
		return conn.ExecCtx(ctx, query, num, pid)
	}, productIdKey)

	return err
}
