package service

import (
	"time"

	"github.com/raystack/salt/log"
	"github.com/robfig/cron/v3"
	"golang.org/x/net/context"

	"github.com/raystack/optimus/config"
	"github.com/raystack/optimus/core/scheduler"
)

const (
	syncInterval = "@every 1m"
)

type ReplayManager struct {
	l log.Logger

	replayRepository ReplayRepository
	replayWorker     Worker

	schedule *cron.Cron
	Now      func() time.Time

	config config.ReplayConfig
}

func NewReplayManager(l log.Logger, replayRepository ReplayRepository, replayWorker Worker, now func() time.Time, config config.ReplayConfig) *ReplayManager {
	_ = "STUB: not implemented"
	return nil
}

type Worker interface {
	Process(*scheduler.ReplayWithRun)
}

func (m ReplayManager) Initialize() { _ = "STUB: not implemented"; return }

func (m ReplayManager) StartReplayLoop() { _ = "STUB: not implemented"; return }

// Cancel timed out replay with status [created, in progress, partial replayed, replayed]

// Fetch created, in progress, and replayed request

func (m ReplayManager) checkTimedOutReplay(ctx context.Context) { _ = "STUB: not implemented"; return }
