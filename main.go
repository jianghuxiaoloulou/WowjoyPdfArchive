package main

import (
	"WowjoyProject/WowjoyPdfArchive/global"
	"WowjoyProject/WowjoyPdfArchive/internal/routers"
	"WowjoyProject/WowjoyPdfArchive/internal/routers/api/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title PDF归档服务
// @version 1.0.0.0
// @description PDF归档服务
// @termsOfService
func main() {
	global.Logger.Info("***开始运行PDF归档服务***")
	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	web()
}

// http
func web() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	ser := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	ser.ListenAndServe()
}
