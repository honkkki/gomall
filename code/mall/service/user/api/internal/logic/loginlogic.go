package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/common/jwt"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/types/user"
	"time"

	"github.com/honkkki/gomall/code/mall/service/user/api/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpcClient.Login(l.ctx, &user.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	expire := now + l.svcCtx.Config.Auth.AccessExpire

	token, err := jwt.GenerateToken(l.svcCtx.Config.Auth.AccessSecret, now, expire, res.Id)

	return &types.LoginResponse{
		AccessToken:  token,
		AccessExpire: expire,
	}, nil
}
