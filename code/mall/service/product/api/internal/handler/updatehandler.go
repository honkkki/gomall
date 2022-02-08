package handler

import (
	"github.com/honkkki/gomall/code/mall/common/retdata"
	"net/http"

	"github.com/honkkki/gomall/code/mall/service/product/api/internal/logic"
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func UpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateLogic(r.Context(), ctx)
		resp, err := l.Update(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, retdata.NewSuccessRet(resp))
		}
	}
}
