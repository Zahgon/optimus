package survey

import (
	"github.com/raystack/salt/log"
)

// BackupCreateSurvey defines survey for creating backup
type BackupCreateSurvey struct {
	logger log.Logger
}

// NewBackupCreateSurvey initializes surveys for creating backup
func NewBackupCreateSurvey(logger log.Logger) *BackupCreateSurvey {
	_ = "STUB: not implemented"
	return nil
}

// AskResourceNames asks the user to add resource name for creating backup
func (*BackupCreateSurvey) AskResourceNames() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AskBackupDescription asks the user the need of backup creation
func (*BackupCreateSurvey) AskBackupDescription() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
