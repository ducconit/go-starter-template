package middleware

import (
	"app/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.Api().AllowOrigins
	corsConfig.AllowAllOrigins = false
	corsConfig.AllowCredentials = config.Api().AllowCredentials
	corsConfig.AddAllowHeaders("Authorization")

	return cors.New(corsConfig)
}
