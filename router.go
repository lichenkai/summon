package summon

import (
	"context"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	*httprouter.Router
	maxBodyBytes int64
	ctx          context.Context
}

func NewRouter(ctx context.Context) *Router {
	r := &Router{
		Router: httprouter.New(),
		ctx:    ctx,
	}
	r.Router.RedirectTrailingSlash = false
	r.Router.RedirectFixedPath = false
	return r
}

func (r *Router) SetMaxBodyBytes(n int64) {
	r.maxBodyBytes = n
}

func (r *Router) Handle(method, pattern string, h ContextHandler) {
	r.Router.Handle(method, pattern, Handle(r.ctx, h, r.maxBodyBytes))
}

func (r *Router) HandleFunc(method, pattern string, h func(context.Context, ResponseWriter, *Request)) {
	r.Handle(method, pattern, ContextHandlerFunc(h))
}

func (r *Router) HEAD(pattern string, h ContextHandler) {
	r.Handle("HEAD", pattern, h)
}

func (r *Router) OPTIONS(pattern string, h ContextHandler) {
	r.Handle("OPTIONS", pattern, h)
}

func (r *Router) GET(pattern string, h ContextHandler) {
	r.Handle("GET", pattern, h)
}

func (r *Router) POST(pattern string, h ContextHandler) {
	r.Handle("POST", pattern, h)
}

func (r *Router) PUT(pattern string, h ContextHandler) {
	r.Handle("PUT", pattern, h)
}

func (r *Router) DELETE(pattern string, h ContextHandler) {
	r.Handle("DELETE", pattern, h)
}

func (r *Router) PATCH(pattern string, h ContextHandler) {
	r.Handle("PATCH", pattern, h)
}
