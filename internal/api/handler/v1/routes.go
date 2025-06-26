package v1

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes cấu hình các routes cho API v1
func SetupRoutes(router *gin.RouterGroup) {
	// Health check toàn cục
	router.GET("/ping", Ping)
	router.GET("/config", GetConfig)
}
