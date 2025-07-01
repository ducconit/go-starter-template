package metrics

import "errors"

var (
	// ErrInvalidMetricType được trả về khi loại metric không hợp lệ
	ErrInvalidMetricType = errors.New("invalid metric type")
	// ErrInconsistentCardinality được trả về khi số lượng labels không khớp
	ErrInconsistentCardinality = errors.New("inconsistent label cardinality")
	// ErrMetricNotFound được trả về khi không tìm thấy metric
	ErrMetricNotFound = errors.New("metric not found")
	// ErrInvalidLabelValues được trả về khi giá trị labels không hợp lệ
	ErrInvalidLabelValues = errors.New("invalid label values")
)
