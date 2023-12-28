package handler

import (
	"github.com/gin-gonic/gin"
)

// InitRouting routesの初期化
func InitRouting(r *gin.Engine, taskHandler TaskHandler) {
	r.POST("/task", taskHandler.Post)
	r.GET("/task/:id", taskHandler.Get)
	r.PUT("/task/:id", taskHandler.Put)
	r.DELETE("/task/:id", taskHandler.Delete)
}
