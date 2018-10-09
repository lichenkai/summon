package handler

import (
	"context"

	"github.com/lichenkai/summon"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(ctx context.Context, f summon.ResponseWriter, r *summon.Request) {
	response := make(map[string]interface{})
	response["id"] = r.PostForm.Get("id")
	response["uid"] = ctx.Value("uid")

	ResponseJson(ctx, f, response)
}
