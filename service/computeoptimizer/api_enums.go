// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computeoptimizer

type FilterName string

// Enum values for FilterName
const (
	FilterNameFinding                  FilterName = "Finding"
	FilterNameRecommendationSourceType FilterName = "RecommendationSourceType"
)

func (enum FilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Finding string

// Enum values for Finding
const (
	FindingUnderprovisioned Finding = "Underprovisioned"
	FindingOverprovisioned  Finding = "Overprovisioned"
	FindingOptimized        Finding = "Optimized"
	FindingNotOptimized     Finding = "NotOptimized"
)

func (enum Finding) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Finding) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MetricName string

// Enum values for MetricName
const (
	MetricNameCpu    MetricName = "Cpu"
	MetricNameMemory MetricName = "Memory"
)

func (enum MetricName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MetricName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MetricStatistic string

// Enum values for MetricStatistic
const (
	MetricStatisticMaximum MetricStatistic = "Maximum"
	MetricStatisticAverage MetricStatistic = "Average"
)

func (enum MetricStatistic) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MetricStatistic) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RecommendationSourceType string

// Enum values for RecommendationSourceType
const (
	RecommendationSourceTypeEc2instance      RecommendationSourceType = "Ec2Instance"
	RecommendationSourceTypeAutoScalingGroup RecommendationSourceType = "AutoScalingGroup"
)

func (enum RecommendationSourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RecommendationSourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Status string

// Enum values for Status
const (
	StatusActive   Status = "Active"
	StatusInactive Status = "Inactive"
	StatusPending  Status = "Pending"
	StatusFailed   Status = "Failed"
)

func (enum Status) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Status) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}