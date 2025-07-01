package metrics

import (
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

// Collector chứa tất cả các metrics của hệ thống
type Collector struct {
	// System metrics
	CPUUtilization *prometheus.GaugeVec
	MemoryUsage    *prometheus.GaugeVec
	DiskUsage      *prometheus.GaugeVec
	GoRoutines     prometheus.Gauge
	GoMemoryAlloc  prometheus.Gauge
	GoTotalAlloc   prometheus.Gauge
	GoSys          prometheus.Gauge
	GoNumGC        prometheus.Counter

	// Application metrics
	HTTPRequestsTotal *prometheus.CounterVec
	HTTPDuration      *prometheus.HistogramVec
	DBConnections     *prometheus.GaugeVec
	DBQueriesTotal    *prometheus.CounterVec

	// Dynamic metrics registry
	customMetrics map[string]any
}

// NewCollector khởi tạo một collector mới
func NewCollector(namespace string) *Collector {
	return &Collector{
		// System metrics
		CPUUtilization: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "cpu_utilization_percent",
			Help:      "Current CPU utilization in percent",
		}, []string{"mode"}),

		MemoryUsage: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "memory_usage_bytes",
			Help:      "Current memory usage in bytes",
		}, []string{"type"}),

		DiskUsage: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "disk_usage_bytes",
			Help:      "Current disk usage in bytes",
		}, []string{"device", "mountpoint", "fstype"}),

		GoRoutines: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "go_goroutines",
			Help:      "Number of goroutines",
		}),

		GoMemoryAlloc: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "go_memstats_alloc_bytes",
			Help:      "Number of bytes allocated and still in use",
		}),

		GoTotalAlloc: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "go_memstats_total_alloc_bytes",
			Help:      "Total number of bytes allocated, even if freed",
		}),

		GoSys: promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "go_memstats_sys_bytes",
			Help:      "Number of bytes obtained from system",
		}),

		GoNumGC: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "go_gc_cycles_total",
			Help:      "Total number of GC cycles",
		}),

		// Application metrics
		HTTPRequestsTotal: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_requests_total",
			Help:      "Total number of HTTP requests",
		}, []string{"method", "path", "status"}),

		HTTPDuration: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "http_request_duration_seconds",
			Help:      "Duration of HTTP requests in seconds",
			Buckets:   []float64{0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		}, []string{"method", "path"}),

		DBConnections: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "db_connections",
			Help:      "Number of database connections",
		}, []string{"db_name", "type"}),

		DBQueriesTotal: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "db_queries_total",
			Help:      "Total number of database queries",
		}, []string{"db_name", "operation", "table"}),

		customMetrics: make(map[string]any),
	}
}

// StartCollecting bắt đầu thu thập metrics hệ thống
func (c *Collector) StartCollecting(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.collectSystemMetrics()
		}
	}()
}

// collectSystemMetrics thu thập các metrics hệ thống
func (c *Collector) collectSystemMetrics() {
	// CPU usage
	if cpuPercents, err := cpu.Percent(0, false); err == nil && len(cpuPercents) > 0 {
		c.CPUUtilization.WithLabelValues("user").Set(cpuPercents[0])
	}

	// Memory usage
	if memInfo, err := mem.VirtualMemory(); err == nil {
		c.MemoryUsage.WithLabelValues("used").Set(float64(memInfo.Used))
		c.MemoryUsage.WithLabelValues("total").Set(float64(memInfo.Total))
		c.MemoryUsage.WithLabelValues("free").Set(float64(memInfo.Free))
	}

	// Disk usage
	if partitions, err := disk.Partitions(true); err == nil {
		for _, p := range partitions {
			if usage, err := disk.Usage(p.Mountpoint); err == nil {
				c.DiskUsage.WithLabelValues(p.Device, p.Mountpoint, p.Fstype).Set(float64(usage.Used))
			}
		}
	}

	// Go runtime metrics
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	c.GoRoutines.Set(float64(runtime.NumGoroutine()))
	c.GoMemoryAlloc.Set(float64(m.Alloc))
	c.GoTotalAlloc.Set(float64(m.TotalAlloc))
	c.GoSys.Set(float64(m.Sys))
}

// RegisterCustomMetric đăng ký một custom metric
func (c *Collector) RegisterCustomMetric(name string, metric any) {
	c.customMetrics[name] = metric
}

// GetCustomMetric lấy một custom metric đã đăng ký
func (c *Collector) GetCustomMetric(name string) (any, bool) {
	metric, exists := c.customMetrics[name]
	return metric, exists
}
