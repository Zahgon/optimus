package telemetry

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	counterMetricMap   = map[string]prometheus.Counter{}
	counterMetricMutex = sync.Mutex{}

	gaugeMetricMap   = map[string]prometheus.Gauge{}
	gaugeMetricMutex = sync.Mutex{}
)

func getKey(metric string, labels map[string]string) string { _ = "STUB: not implemented"; return "" }

func NewCounter(metric string, labels map[string]string) prometheus.Counter {
	_ = "STUB: not implemented"
	return *new(prometheus.Counter)
}

func NewGauge(metric string, labels map[string]string) prometheus.Gauge {
	_ = "STUB: not implemented"
	return *new(prometheus.Gauge)
}
