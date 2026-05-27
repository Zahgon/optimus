package airflow

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/raystack/optimus/core/scheduler"
	"github.com/raystack/optimus/internal/lib/cron"
)

const (
	pageLimit = 99999
)

type airflowRequest struct {
	path   string
	method string
	body   []byte
}

type DagRunListResponse struct {
	DagRuns      []DagRun `json:"dag_runs"`
	TotalEntries int      `json:"total_entries"`
}

type DagRun struct {
	ExecutionDate   time.Time `json:"execution_date"`
	State           string    `json:"state"`
	ExternalTrigger bool      `json:"external_trigger"`
}

type DagRunRequest struct {
	OrderBy          string   `json:"order_by"`
	PageOffset       int      `json:"page_offset"`
	PageLimit        int      `json:"page_limit"`
	DagIds           []string `json:"dag_ids"`
	ExecutionDateGte string   `json:"execution_date_gte,omitempty"`
	ExecutionDateLte string   `json:"execution_date_lte,omitempty"`
}

type SchedulerAuth struct {
	host  string
	token string
}

type ClientAirflow struct {
	client *http.Client
}

func NewAirflowClient() *ClientAirflow { _ = "STUB: not implemented"; return nil }

func (ac ClientAirflow) Invoke(ctx context.Context, r airflowRequest, auth SchedulerAuth) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseResponse(resp *http.Response) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func buildEndPoint(host, path string) string { _ = "STUB: not implemented"; return "" }

func getJobRuns(res DagRunListResponse, spec *cron.ScheduleSpec) ([]*scheduler.JobRunStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only include scheduled runs

// use multi error to collect errors and proceed

func startChildSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}
