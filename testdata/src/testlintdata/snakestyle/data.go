package snakestyle

type MetricData map[string]any

func literalValid() {
	_ = MetricData{
		"request_count": 1,
		"error_rate":    0.01,
	}
}

func literalInvalid() {
	_ = MetricData{
		"request_count": 1,
		"ErrorRate":     0.01, // want `Key 'ErrorRate' style should be snake_case`
	}
}

const (
	validSnakeConst   = "response_time"
	invalidSnakeConst = "ResponseTime"
)

func constKeyValid() {
	_ = MetricData{validSnakeConst: 42}
}

func constKeyInvalid() {
	_ = MetricData{invalidSnakeConst: 42} // want `Key 'ResponseTime' \(from identifier 'invalidSnakeConst'\) style should be snake_case`
}
