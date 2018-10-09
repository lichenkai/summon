package main

import (
	"log"
	"os"

	"github.com/lichenkai/summon/demo/handler"
	"github.com/lichenkai/summon/grace"
)

func main() {
	h := handler.Register()
	err := grace.ListenAndServe("0.0.0.0:9999", h)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 9999 stopped")
	os.Exit(0)
}
