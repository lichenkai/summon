package middleware

import (
	"context"
	"net/http"

	"github.com/lichenkai/summon"
)

func Recovery(h summon.ContextHandler) summon.ContextHandler {
	f := func(ctx context.Context, w summon.ResponseWriter, r *summon.Request) {
		defer func() {
			if r := recover(); r != nil {
				switch r.(type) {
				default:
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
				}
			}
		}()

		h.ServeHTTP(ctx, w, r)
	}
	return summon.ContextHandlerFunc(f)
}
