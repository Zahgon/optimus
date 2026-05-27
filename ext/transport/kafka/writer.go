package kafka

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/raystack/salt/log"
	"github.com/segmentio/kafka-go"
)

const (
	writeTimeout = time.Second * 3
)

var kafkaQueueCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "publisher_kafka_events_queued_total",
	Help: "Number of events queued to be published to kafka topic",
})

type Writer struct {
	kafkaWriter *kafka.Writer
}

func NewWriter(kafkaBrokerUrls []string, topic string, logger log.Logger) *Writer {
	_ = "STUB: not implemented"
	return nil
}

func (w *Writer) Close() error { _ = "STUB: not implemented"; return nil }

func (w *Writer) Write(messages [][]byte) error { _ = "STUB: not implemented"; return nil }
