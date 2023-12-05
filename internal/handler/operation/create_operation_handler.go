package operation

import (
	"net/http"
	"server/response"

	"server/internal/logic/operation"
	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateOperationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOperationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := operation.NewCreateOperationLogic(r.Context(), svcCtx)
		resp, err := l.CreateOperation(&req)
		response.Response(w, resp, err)
	}
}
