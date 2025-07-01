package monitor

import (
	"app/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter(router *gin.Engine) {
	if config.EnableMetrics() {
		// Thêm endpoint metrics cho Prometheus
		router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
}
