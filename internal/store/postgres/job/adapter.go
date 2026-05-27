package job

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"

	"github.com/raystack/optimus/core/job"
	"github.com/raystack/optimus/internal/models"
)

const jobDatetimeLayout = "2006-01-02"

type Spec struct {
	ID          uuid.UUID
	Name        string
	Version     int
	Owner       string
	Description string
	Labels      map[string]string

	Schedule   json.RawMessage
	WindowSpec json.RawMessage

	Alert json.RawMessage

	StaticUpstreams pq.StringArray
	HTTPUpstreams   json.RawMessage

	TaskName   string
	TaskConfig map[string]string

	Hooks json.RawMessage

	Assets map[string]string

	Metadata json.RawMessage

	Destination string
	Sources     pq.StringArray

	ProjectName   string `json:"project_name"`
	NamespaceName string `json:"namespace_name"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type Schedule struct {
	StartDate     time.Time
	EndDate       *time.Time `json:",omitempty"`
	Interval      string
	DependsOnPast bool
	Retry         *Retry
}

type Window struct {
	WindowSize       string
	WindowOffset     string
	WindowTruncateTo string
}

type Retry struct {
	Count              int
	Delay              int32
	ExponentialBackoff bool
}

type Alert struct {
	On       string
	Config   map[string]string
	Channels []string
}

type Asset struct {
	Name  string
	Value string
}

type Hook struct {
	Name   string
	Config map[string]string
}

type Metadata struct {
	Resource  *MetadataResource
	Scheduler map[string]string
}

type MetadataResource struct {
	Request *MetadataResourceConfig
	Limit   *MetadataResourceConfig
}

type MetadataResourceConfig struct {
	CPU    string
	Memory string
}

type Config struct {
	Configs map[string]string
}

func toStorageSpec(jobEntity *job.Job) (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func toStorageWindow(windowSpec models.Window) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toStorageHooks(hookSpecs []*job.Hook) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toStorageHook(spec *job.Hook) Hook { _ = "STUB: not implemented"; return *new(Hook) }

func toStorageAlerts(alertSpecs []*job.AlertSpec) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toStorageSchedule(scheduleSpec *job.Schedule) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toStorageMetadata(metadataSpec *job.Metadata) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromStorageSpec(jobSpec *Spec) (*job.Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func fromStorageWindow(raw []byte, jobVersion int) (models.Window, error) {
	_ = "STUB: not implemented"
	return *new(models.Window), nil
}

func fromStorageSchedule(raw []byte) (*job.Schedule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromStorageHooks(raw []byte) ([]*job.Hook, error) { _ = "STUB: not implemented"; return nil, nil }

func fromStorageHook(hook Hook) (*job.Hook, error) { _ = "STUB: not implemented"; return nil, nil }

func fromStorageAlerts(raw []byte) ([]*job.AlertSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FromRow(row pgx.Row) (*Spec, error) { _ = "STUB: not implemented"; return nil, nil }

func UpstreamFromRow(row pgx.Row) (*JobWithUpstream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
