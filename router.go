package main

import (
	"eshop_im/handler"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	apiGroup := r.Group("/api/v1")

	// 新增IM路由
	imGroup := apiGroup.Group("/im")
	imGroup.GET("/ws", handler.HandleWebSocket)
	imGroup.GET("/receiver/mget", handler.HandleMgetReceiver)
	imGroup.POST("/history/get_one", handler.HandleOneHistory)
	//imGroup.GET("/online_users", handler.HandleOnlineUsers)
}
