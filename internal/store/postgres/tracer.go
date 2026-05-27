package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// tracer is a wrapper around the pgx tracer interfaces which instrument
// queries.
type tracer struct {
	tracer trace.Tracer
	attrs  []attribute.KeyValue
}

// NewTracer returns a new Tracer.
func newTracer() *tracer { _ = "STUB: not implemented"; return nil }

func recordError(span trace.Span, err error) { _ = "STUB: not implemented"; return }

// TraceQueryStart is called at the beginning of Query, QueryRow, and Exec calls.
// The returned context is used for the rest of the call and will be passed to TraceQueryEnd.
func (t *tracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// TraceQueryEnd is called at the end of Query, QueryRow, and Exec calls.
func (*tracer) TraceQueryEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryEndData) {
	_ = "STUB: not implemented"
	return
}

// TraceCopyFromStart is called at the beginning of CopyFrom calls. The
// returned context is used for the rest of the call and will be passed to
// TraceCopyFromEnd.
func (t *tracer) TraceCopyFromStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceCopyFromStartData) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// TraceCopyFromEnd is called at the end of CopyFrom calls.
func (*tracer) TraceCopyFromEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceCopyFromEndData) {
	_ = "STUB: not implemented"
	return
}

// TraceBatchStart is called at the beginning of SendBatch calls. The returned
// context is used for the rest of the call and will be passed to
// TraceBatchQuery and TraceBatchEnd.
func (t *tracer) TraceBatchStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceBatchStartData) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// TraceBatchQuery is called at the after each query in a batch.
func (t *tracer) TraceBatchQuery(ctx context.Context, _ *pgx.Conn, data pgx.TraceBatchQueryData) {
	_ = "STUB: not implemented"
	return
}

// TraceBatchEnd is called at the end of SendBatch calls.
func (*tracer) TraceBatchEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceBatchEndData) {
	_ = "STUB: not implemented"
	return
}

// TraceConnectStart is called at the beginning of Connect and ConnectConfig
// calls. The returned context is used for the rest of the call and will be
// passed to TraceConnectEnd.
func (t *tracer) TraceConnectStart(ctx context.Context, data pgx.TraceConnectStartData) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// TraceConnectEnd is called at the end of Connect and ConnectConfig calls.
func (*tracer) TraceConnectEnd(ctx context.Context, data pgx.TraceConnectEndData) {
	_ = "STUB: not implemented"
	return
}

// TracePrepareStart is called at the beginning of Prepare calls. The returned
// context is used for the rest of the call and will be passed to
// TracePrepareEnd.
func (t *tracer) TracePrepareStart(ctx context.Context, _ *pgx.Conn, data pgx.TracePrepareStartData) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// TracePrepareEnd is called at the end of Prepare calls.
func (*tracer) TracePrepareEnd(ctx context.Context, _ *pgx.Conn, data pgx.TracePrepareEndData) {
	_ = "STUB: not implemented"
	return
}
