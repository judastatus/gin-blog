package main

import (
	"context"
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/gredis"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @license.name MIT
func main() {

	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()

	// *************** 1 不支持重启之前 ***********
	//router := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//s.ListenAndServe()

	// *************** 2 endless重启的方式 ***********
	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}

	// *************** 3 http.Server 的 Shutdown 方法来重启 ***********
	//router := routers.InitRouter()
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//go func() {
	//	if err := s.ListenAndServe(); err != nil {
	//		log.Printf("Listen: %s\n", err)
	//	}
	//}()
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<- quit
	//log.Println("Shutdown Server ...")
	//ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	//defer cancel()
	//if err := s.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
	//log.Println("Server exiting")


	// *************** 4 http.Server 的 Shutdown 方法来重启 ***********

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	s := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}