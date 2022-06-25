package core

import (
	"gin-web/global"
	"gin-web/initialiaze"
	"gin-web/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
	ListenAndServeTLS(certFile string, keyFile string) error
}

func RunServer() {
	Router := initialiaze.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Port)
	s := initServer(address, Router)
	time.Sleep(time.Millisecond * 100)
	protocol := "http://"
	fmt.Printf("############################################################################\n")
	fmt.Printf("\n        Project Start Success")
	fmt.Printf("\n        Server Listen at:")
	for _, ip := range utils.GetLocalIP() {
		fmt.Printf("\n        %s%s%s", protocol, ip, address)
	}
	fmt.Printf("\n\n############################################################################\n")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
