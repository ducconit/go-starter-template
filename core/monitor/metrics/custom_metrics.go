package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// CustomMetricType định nghĩa các loại metric tùy chỉnh
const (
	GaugeType     = "gauge"
	CounterType   = "counter"
	HistogramType = "histogram"
	SummaryType   = "summary"
)

// CustomMetricConfig cấu hình cho metric tùy chỉnh
type CustomMetricConfig struct {
	Name       string
	Help       string
	Type       string
	Labels     []string
	Buckets    []float64           // Chỉ dùng cho Histogram
	Objectives map[float64]float64 // Chỉ dùng cho Summary
}

// CreateCustomMetric tạo một metric tùy chỉnh dựa trên cấu hình
func (c *Collector) CreateCustomMetric(config CustomMetricConfig) (any, error) {
	switch config.Type {
	case GaugeType:
		metric := promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: config.Name,
			Help: config.Help,
		}, config.Labels)
		c.customMetrics[config.Name] = metric
		return metric, nil

	case CounterType:
		metric := promauto.NewCounterVec(prometheus.CounterOpts{
			Name: config.Name,
			Help: config.Help,
		}, config.Labels)
		c.customMetrics[config.Name] = metric
		return metric, nil

	case HistogramType:
		buckets := config.Buckets
		if len(buckets) == 0 {
			buckets = prometheus.DefBuckets
		}
		metric := promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name:    config.Name,
			Help:    config.Help,
			Buckets: buckets,
		}, config.Labels)
		c.customMetrics[config.Name] = metric
		return metric, nil

	case SummaryType:
		objectives := config.Objectives
		if len(objectives) == 0 {
			objectives = map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}
		}
		metric := promauto.NewSummaryVec(prometheus.SummaryOpts{
			Name:       config.Name,
			Help:       config.Help,
			Objectives: objectives,
		}, config.Labels)
		c.customMetrics[config.Name] = metric
		return metric, nil

	default:
		return nil, ErrInvalidMetricType
	}
}

// GetOrCreateGauge lấy hoặc tạo mới một Gauge metric
func (c *Collector) GetOrCreateGauge(name, help string, labels []string) (*prometheus.GaugeVec, error) {
	if metric, exists := c.customMetrics[name]; exists {
		if gauge, ok := metric.(*prometheus.GaugeVec); ok {
			return gauge, nil
		}
		return nil, ErrInconsistentCardinality
	}

	gauge := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	}, labels)

	c.customMetrics[name] = gauge
	return gauge, nil
}

// GetOrCreateCounter lấy hoặc tạo mới một Counter metric
func (c *Collector) GetOrCreateCounter(name, help string, labels []string) (*prometheus.CounterVec, error) {
	if metric, exists := c.customMetrics[name]; exists {
		if counter, ok := metric.(*prometheus.CounterVec); ok {
			return counter, nil
		}
		return nil, ErrInconsistentCardinality
	}

	counter := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: name,
		Help: help,
	}, labels)

	c.customMetrics[name] = counter
	return counter, nil
}

// GetOrCreateHistogram lấy hoặc tạo mới một Histogram metric
func (c *Collector) GetOrCreateHistogram(name, help string, buckets []float64, labels []string) (*prometheus.HistogramVec, error) {
	if metric, exists := c.customMetrics[name]; exists {
		if hist, ok := metric.(*prometheus.HistogramVec); ok {
			return hist, nil
		}
		return nil, ErrInconsistentCardinality
	}

	if len(buckets) == 0 {
		buckets = prometheus.DefBuckets
	}

	hist := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    name,
		Help:    help,
		Buckets: buckets,
	}, labels)

	c.customMetrics[name] = hist
	return hist, nil
}

// GetOrCreateSummary lấy hoặc tạo mới một Summary metric
func (c *Collector) GetOrCreateSummary(name, help string, objectives map[float64]float64, labels []string) (*prometheus.SummaryVec, error) {
	if metric, exists := c.customMetrics[name]; exists {
		if summary, ok := metric.(*prometheus.SummaryVec); ok {
			return summary, nil
		}
		return nil, ErrInconsistentCardinality
	}

	if len(objectives) == 0 {
		objectives = map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}
	}

	summary := promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name:       name,
		Help:       help,
		Objectives: objectives,
	}, labels)

	c.customMetrics[name] = summary
	return summary, nil
}
