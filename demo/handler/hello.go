package handler

import (
	"context"
	"fmt"

	"github.com/lichenkai/summon"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(ctx context.Context, f summon.ResponseWriter, r *summon.Request) {
	fmt.Println(r.Form.Get("id"))
}
