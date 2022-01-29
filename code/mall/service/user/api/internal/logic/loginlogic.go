package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/common/jwt"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/userclient"
	"time"

	"github.com/honkkki/gomall/code/mall/service/user/api/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	expSeconds := l.svcCtx.Config.Auth.AccessExpire

	token, err := jwt.GenerateToken(l.svcCtx.Config.Auth.AccessSecret, now, expSeconds, res.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		AccessToken:  token,
		AccessExpire: now + expSeconds, // token到期时间戳
	}, nil
}
