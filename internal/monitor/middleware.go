package monitor

import (
	"app/internal/config"

	"github.com/gin-gonic/gin"
)

func HTTPMetricsMiddleware() gin.HandlerFunc {
	if metricsCollector == nil || !config.EnableMetrics() {
		return func(ctx *gin.Context) {
			ctx.Next()
		}
	}
	return metricsCollector.HTTPMetricsMiddleware()
}
