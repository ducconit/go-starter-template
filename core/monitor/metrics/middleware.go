package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// HTTPMetricsMiddleware tạo middleware để ghi nhận HTTP metrics
func (c *Collector) HTTPMetricsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bỏ qua các endpoint health check
		if ctx.Request.URL.Path == "/health" || ctx.Request.URL.Path == "/metrics" {
			ctx.Next()
			return
		}

		start := time.Now()

		// Xử lý request
		ctx.Next()

		// Tính thời gian xử lý
		duration := time.Since(start).Seconds()

		// Lấy thông tin request
		status := strconv.Itoa(ctx.Writer.Status())
		method := ctx.Request.Method
		path := ctx.FullPath()

		// Nếu không tìm thấy path (404), gán là "not_found"
		if path == "" {
			path = "not_found"
		}

		// Ghi nhận metrics
		c.HTTPRequestsTotal.WithLabelValues(method, path, status).Inc()
		c.HTTPDuration.WithLabelValues(method, path).Observe(duration)
	}
}

// DatabaseMetricsMiddleware tạo middleware để ghi nhận database metrics
func (c *Collector) DatabaseMetricsMiddleware(dbName, operation, table string) func() {
	// Tăng counter cho số lượng query
	c.DBQueriesTotal.WithLabelValues(dbName, operation, table).Inc()

	// Trả về hàm để cập nhật số lượng connection khi hoàn thành
	return func() {
		// Có thể thêm logic để cập nhật số lượng connection nếu cần
	}
}

// UpdateDBConnections cập nhật số lượng kết nối database
func (c *Collector) UpdateDBConnections(dbName, connType string, count int) {
	c.DBConnections.WithLabelValues(dbName, connType).Set(float64(count))
}
