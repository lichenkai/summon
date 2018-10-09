package handler

import (
	"context"

	"github.com/lichenkai/summon"
)

func Login(h summon.ContextHandler) summon.ContextHandler {
	f := func(ctx context.Context, w summon.ResponseWriter, r *summon.Request) {
		ctx = context.WithValue(ctx, "uid", 123)
		h.ServeHTTP(ctx, w, r)
	}
	return summon.ContextHandlerFunc(f)
}
