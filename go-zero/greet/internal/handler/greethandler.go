package handler

import (
	"net/http"

	"github.com/feixiao/learning/go-zero/greet/internal/logic"
	"github.com/feixiao/learning/go-zero/greet/internal/svc"
	"github.com/feixiao/learning/go-zero/greet/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func greetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), ctx)
		resp, err := l.Greet(req)

		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
