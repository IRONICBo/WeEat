package handler

import (
	"net/http"

	"api-service/internal/logic"
	"api-service/internal/svc"
	"api-service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RAGSyncChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RAGChatMessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRAGSyncChatLogic(r.Context(), svcCtx)
		resp, err := l.RAGSyncChat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}