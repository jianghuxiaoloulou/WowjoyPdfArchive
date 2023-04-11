package routers

import (
	v1 "WowjoyProject/WowjoyPdfArchive/internal/routers/api/v1"
	"WowjoyProject/WowjoyPdfArchive/internal/routers/api/ws"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 注册中间件
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		// 通过数据流上传文件
		apiv1.POST("/SaveFile", v1.ByFileStream)
	}
	wsGroup := r.Group("/api/ws")
	{
		// websocket
		wsGroup.GET("/:uidenc", ws.WebsocketManager.WsClient)
	}
	return r
}
