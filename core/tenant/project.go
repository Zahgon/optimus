package tenant

const (
	EntityProject = "project"

	ProjectStoragePathKey = "STORAGE_PATH"
	ProjectSchedulerHost  = "SCHEDULER_HOST"
)

type ProjectName string

func ProjectNameFrom(name string) (ProjectName, error) {
	_ = "STUB: not implemented"
	return *new(ProjectName), nil
}

func (pn ProjectName) String() string { _ = "STUB: not implemented"; return "" }

type Project struct {
	name   ProjectName
	config map[string]string
}

func (p *Project) Name() ProjectName { _ = "STUB: not implemented"; return *new(ProjectName) }

func (p *Project) GetConfig(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetConfigs returns a clone of project configurations
func (p *Project) GetConfigs() map[string]string { _ = "STUB: not implemented"; return nil }

func NewProject(name string, config map[string]string) (*Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
