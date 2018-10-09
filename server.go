package summon

import (
	"fmt"
	"net/http"
)

func GetHttpServer() *http.Server {

	// 注册router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	// 定义server对象
	server := &http.Server{
		Addr:           "",
		ReadTimeout:    1,
		WriteTimeout:   1,
		MaxHeaderBytes: 102400,
	}

	return server
}
