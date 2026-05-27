package tenant

const EntityNamespace = "namespace"

type NamespaceName string

func NamespaceNameFrom(name string) (NamespaceName, error) {
	_ = "STUB: not implemented"
	return *new(NamespaceName), nil
}

func (n NamespaceName) String() string { _ = "STUB: not implemented"; return "" }

type Namespace struct {
	name NamespaceName

	projectName ProjectName
	config      map[string]string
}

func (n *Namespace) Name() NamespaceName { _ = "STUB: not implemented"; return *new(NamespaceName) }

func (n *Namespace) ProjectName() ProjectName { _ = "STUB: not implemented"; return *new(ProjectName) }

func (n *Namespace) GetConfig(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetConfigs returns a clone on project configurations
func (n *Namespace) GetConfigs() map[string]string { _ = "STUB: not implemented"; return nil }

func NewNamespace(name string, projName ProjectName, config map[string]string) (*Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
