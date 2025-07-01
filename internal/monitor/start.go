package monitor

import (
	"app/internal/config"
	"core/monitor/metrics"
	"time"
)

var metricsCollector *metrics.Collector

func Start(appName string, interval time.Duration) {
	if !config.EnableMetrics() {
		return
	}
	// Khởi tạo metrics collector
	metricsCollector = metrics.NewCollector(appName)
	// Bắt đầu thu thập metrics hệ thống
	metricsCollector.StartCollecting(interval)
}
