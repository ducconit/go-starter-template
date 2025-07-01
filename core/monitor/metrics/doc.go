// Package metrics cung cấp các công cụ để thu thập và xuất metrics của hệ thống
// và ứng dụng sử dụng Prometheus.
//
// Các tính năng chính:
// - Thu thập metrics hệ thống: CPU, RAM, Disk, Goroutines, v.v.
// - Metrics ứng dụng: HTTP requests, database queries, v.v.
// - Hỗ trợ tạo và quản lý custom metrics linh hoạt
// - Middleware cho Gin framework để tự động thu thập HTTP metrics
//
// Ví dụ sử dụng:
/*
package main

import (
	"time"
	"your-project/core/monitor/metrics"
	"github.com/gin-gonic/gin"
)

func main() {
	// Khởi tạo collector với namespace cho ứng dụng
	collector := metrics.NewCollector("myapp")

	// Bắt đầu thu thập metrics hệ thống mỗi 5 giây
	collector.StartCollecting(5 * time.Second)

	// Tạo một HTTP server với Gin
	r := gin.Default()

	// Sử dụng middleware để thu thập HTTP metrics
	r.Use(collector.HTTPMetricsMiddleware())

	// Thêm endpoint metrics cho Prometheus
	// Lưu ý: Cần import "github.com/prometheus/client_golang/prometheus/promhttp"
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Thêm một endpoint ví dụ
	r.GET("/api/example", func(c *gin.Context) {
		// Tăng counter cho database query
		defer collector.DBQueriesTotal.WithLabelValues("main_db", "select", "users").Inc()

		// Cập nhật số lượng kết nối database
		collector.UpdateDBConnections("main_db", "active", 10)

		c.JSON(200, gin.H{"message": "success"})
	})

	// Tạo một custom metric
	customMetric, _ := collector.GetOrCreateGauge(
		"custom_metric_example",
		"Một custom metric ví dụ",
		[]string{"label1", "label2"},
	)

	// Sử dụng custom metric
	customMetric.WithLabelValues("value1", "value2").Set(42)

	r.Run(":8080")
}
*/
package metrics
