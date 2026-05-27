package config

type ClientConfig struct {
	Version    Version      `mapstructure:"version"`
	Log        LogConfig    `mapstructure:"log"`
	Host       string       `mapstructure:"host"` // optimus server host
	Project    Project      `mapstructure:"project"`
	Namespaces []*Namespace `mapstructure:"namespaces"`
	Auth       Auth         `mapstructure:"auth"`

	namespaceNameToNamespace map[string]*Namespace
}

type Datastore struct {
	Type   string            `mapstructure:"type"`   // type could be bigquery/postgres/gcs
	Path   string            `mapstructure:"path"`   // directory to find specifications
	Backup map[string]string `mapstructure:"backup"` // backup configuration
}

type Job struct {
	Path string `mapstructure:"path"` // directory to find specifications
}

type Project struct {
	Name   string            `mapstructure:"name"`
	Config map[string]string `mapstructure:"config"`
}

type Auth struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
}

type Namespace struct {
	Name      string            `mapstructure:"name"`
	Config    map[string]string `mapstructure:"config"`
	Job       Job               `mapstructure:"job"`
	Datastore []Datastore       `mapstructure:"datastore"`
}

func (c *ClientConfig) GetNamespaceByName(name string) (*Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClientConfig) ValidateNamespaceNames(namespaceNames ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClientConfig) GetSelectedNamespaces(namespaceNames ...string) ([]*Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClientConfig) GetAllNamespaceNames() []string { _ = "STUB: not implemented"; return nil }

func (c *ClientConfig) buildDictionary() { _ = "STUB: not implemented"; return }
