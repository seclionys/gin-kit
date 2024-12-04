package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

type server interface {
	ListenAndServe() error
}

func initServer(addr string, router *gin.Engine) server {
	server := endless.NewServer(addr, router)
	server.ReadHeaderTimeout = 10 * time.Minute
	server.WriteTimeout = 10 * time.Minute
	server.MaxHeaderBytes = 1 << 20
	return server
}

func RunServer(addr string, router *gin.Engine) {
	server := initServer(addr, router)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
