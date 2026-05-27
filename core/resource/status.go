package resource

type Status string

const (
	StatusUnknown           Status = "unknown"
	StatusValidationFailure Status = "validation_failure"
	StatusValidationSuccess Status = "validation_success"
	StatusToCreate          Status = "to_create"
	StatusToUpdate          Status = "to_update"
	StatusSkipped           Status = "skipped"
	StatusCreateFailure     Status = "create_failure"
	StatusUpdateFailure     Status = "update_failure"
	StatusExistInStore      Status = "exist_in_store"
	StatusSuccess           Status = "success"
)

func (s Status) String() string { _ = "STUB: not implemented"; return "" }

func FromStringToStatus(status string) Status { _ = "STUB: not implemented"; return *new(Status) }

func StatusForToCreate(status Status) bool { _ = "STUB: not implemented"; return false }

func StatusForToUpdate(status Status) bool { _ = "STUB: not implemented"; return false }

func StatusIsSuccess(status Status) bool { _ = "STUB: not implemented"; return false }

type SyncResponse struct {
	ResourceNames    []string
	IgnoredResources []IgnoredResource
}
