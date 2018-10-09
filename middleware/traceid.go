package middleware

import (
	"context"

	"github.com/lichenkai/summon"
	"github.com/lichenkai/summon/utils"
)

func TraceId(h summon.ContextHandler) summon.ContextHandler {
	f := func(ctx context.Context, w summon.ResponseWriter, r *summon.Request) {
		traceId := r.Header.Get("X-Trace-Id")
		if traceId == "" {
			traceId = utils.FastUUIDStr()
		}
		ctx = context.WithValue(ctx, "X-Trace-Id", traceId)
		h.ServeHTTP(ctx, w, r)
	}
	return summon.ContextHandlerFunc(f)
}
