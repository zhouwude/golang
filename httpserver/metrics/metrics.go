package httpservermetrics

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// 在 Go 语言中，你可以省略类型说明符 [type]，因为编译器可以根据常量的值来推断其类型。
const (
	MetricsNamespace = "httpserver"
)

// 大写才可见
func Register() {
	// 把自定义的指标注册进去
	err := prometheus.Register(functionLatency)
	if err != nil {
		fmt.Println(err)
	}

}

var (
	functionLatency = CreateExecutionTimeMetrics(MetricsNamespace, "Time spent")
)

func NewTimer() *ExecutionTimer {
	return NewExecutionTimer(functionLatency)
}
func NewExecutionTimer(histo *prometheus.HistogramVec) *ExecutionTimer {
	now := time.Now()
	return &ExecutionTimer{
		histo: histo,
		start: now,
		end:   now,
	}

}

type ExecutionTimer struct {
	histo *prometheus.HistogramVec
	start time.Time
	end   time.Time
}

// 计算当前时间和 start 的差值
func (t *ExecutionTimer) ObserveTotal() {
	t.histo.WithLabelValues("total-q").Observe(time.Now().Sub(t.start).Seconds())
}

// 创建指标
func CreateExecutionTimeMetrics(namespace string, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: namespace, Name: "execution_latency_seconds", Help: help, Buckets: prometheus.ExponentialBuckets(0.001, 2, 15)},
		[]string{"stepLabel"})
}
