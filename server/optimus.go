package server

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/raystack/salt/log"
	"google.golang.org/grpc"

	"github.com/raystack/optimus/config"
	"github.com/raystack/optimus/core/event/moderator"
	"github.com/raystack/optimus/internal/models"
)

const keyLength = 32

type setupFn func() error

type OptimusServer struct {
	conf   *config.ServerConfig
	logger log.Logger

	dbPool *pgxpool.Pool
	key    *[keyLength]byte

	serverAddr string
	grpcServer *grpc.Server
	httpServer *http.Server

	pluginRepo *models.PluginRepository
	cleanupFn  []func()

	eventHandler moderator.Handler
}

func New(conf *config.ServerConfig) (*OptimusServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *OptimusServer) setupPublisher() error { _ = "STUB: not implemented"; return nil }

func (s *OptimusServer) setupPlugins() error { _ = "STUB: not implemented"; return nil }

// discover and load plugins.

func (s *OptimusServer) setupTelemetry() error { _ = "STUB: not implemented"; return nil }

func (s *OptimusServer) setupAppKey() error { _ = "STUB: not implemented"; return nil }

func applicationKeyFromString(appKey string) (*[keyLength]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *OptimusServer) setupDB() error { _ = "STUB: not implemented"; return nil }

func (s *OptimusServer) setupGRPCServer() error { _ = "STUB: not implemented"; return nil }

func (s *OptimusServer) setupMonitoring() error { _ = "STUB: not implemented"; return nil }

func (s *OptimusServer) setupHTTPProxy() error { _ = "STUB: not implemented"; return nil }

func (s *OptimusServer) startListening() {
	_ = "STUB: not implemented"
	// run our server in a goroutine so that it doesn't block to wait for termination requests
	return
}

func (s *OptimusServer) Shutdown() { _ = "STUB: not implemented"; return }

// Create a deadline to wait for server

// Todo: log all the errors from cleanup before exit

func (s *OptimusServer) setupHandlers() error {
	_ = "STUB: not implemented"
	// Tenant Bounded Context Setup
	return nil
}

// Scheduler bounded context

// Job Bounded Context Setup

// Resource Bounded Context

// Register datastore

// Tenant Handlers

// Resource Handler

// backup service

// version service

// Core Job Handler
