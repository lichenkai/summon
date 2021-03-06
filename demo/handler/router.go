package handler

import (
	"context"

	"github.com/lichenkai/summon"
	"github.com/lichenkai/summon/middleware"
)

var ctx = context.Background()

func Register() *summon.Router {
	// 定义中间件栈，可根据需要在下面追加
	chain := summon.NewChain(
		middleware.TraceId,
		middleware.Recovery,
		Login,
	)

	// 注册Handler
	router := summon.NewRouter(ctx)
	router.SetMaxBodyBytes(102400)
	router.POST("/hello", chain.Then(&HelloHandler{}))

	return router
}
