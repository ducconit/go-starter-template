package api

import (
	v1 "app/internal/api/handler/v1"
	"app/internal/config"

	swaggerDocs "app/docs/swagger"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter setup router
func SetupRouter(router *gin.Engine) {
	r := router.Group("/api")
	{
		// Cấu hình Swagger nếu được bật
		if config.Api().EnableSwagger {
			swaggerDocs.SwaggerInfo.Host = ""
			r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		}

		// Nhóm tất cả các API dưới tiền tố /api/v1
		apiV1 := r.Group("/v1")
		{
			apiV1.GET("/ping", v1.Ping)
			apiV1.GET("/config", v1.GetConfig)
		}
	}
}
