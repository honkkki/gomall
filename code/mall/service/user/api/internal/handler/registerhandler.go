package handler

import (
	"github.com/honkkki/gomall/code/mall/common/retdata"
	"net/http"

	"github.com/honkkki/gomall/code/mall/service/user/api/internal/logic"
	"github.com/honkkki/gomall/code/mall/service/user/api/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/user/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, retdata.NewSuccessRet(resp))
		}
	}
}
