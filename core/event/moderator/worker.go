package moderator

import (
	"context"
	"sync"
	"time"

	"github.com/raystack/salt/log"
)

type Writer interface {
	Write(messages [][]byte) error
	Close() error
}

type Worker struct {
	mu          sync.Mutex
	wg          sync.WaitGroup
	messageChan <-chan []byte
	closeChan   chan bool

	writer        Writer
	batchInterval time.Duration

	messages [][]byte

	logger log.Logger
}

func NewWorker(messageChan <-chan []byte, writer Writer, batchInterval time.Duration, logger log.Logger) *Worker {
	_ = "STUB: not implemented"
	return nil
}

func (w *Worker) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (w *Worker) Flush() { _ = "STUB: not implemented"; return }

func (w *Worker) Close() error { _ = "STUB: not implemented"; return nil }
