package handler

import (
	"context"
	"fmt"

	"github.com/lichenkai/summon"
)

type ResponseJsonData struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func ResponseJson(ctx context.Context, w summon.ResponseWriter, data interface{}) {
	response := &ResponseJsonData{
		Status: 10000,
		Msg:    "success",
		Data:   data,
	}

	if err := w.WriteJSON(response); err != nil {
		fmt.Println("w.WriteJSON fail")
	}
}
