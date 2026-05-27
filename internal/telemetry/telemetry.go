package telemetry

import (
	"net/http"
	"time"

	"github.com/raystack/salt/log"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/raystack/optimus/config"
)

const MetricWaitInterval = time.Second * 2

func Init(l log.Logger, conf config.TelemetryConfig) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Register our TracerProvider as the global so any imported
// instrumentation in the future will default to using it.

// Traces can extend beyond a single process. This requires context propagation, a mechanism where identifiers for a trace are sent to remote processes.
// TextMapPropagator performs the injection and extraction of a cross-cutting concern value as string key/values
// pairs into carriers that travel in-band across process boundaries.
// The carrier of propagated data on both the client (injector) and server (extractor) side is usually an HTTP request.
// In order to increase compatibility, the key/value pairs MUST only consist of US-ASCII characters that make up
// valid HTTP header fields as per RFC 7230.

// custom metric for app uptime

// start exposing metrics

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	_ = "STUB: not implemented"
	// create the Jaeger exporter
	return nil, nil
}

// Always be sure to batch in production

// Record information about this application in an Resource

func MetricsServer(addr string) *http.Server { _ = "STUB: not implemented"; return nil }

//nolint: gosec
