package logic

import (
	"context"

	"github.com/honkkki/gomall/code/mall/service/product/rpc/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *product.RemoveRequest) (*product.RemoveResponse, error) {

	return &product.RemoveResponse{}, nil
}
