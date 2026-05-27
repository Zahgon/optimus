package config

type Optimus struct {
	// configuration version
	Version int `mapstructure:"version"`
	// optimus server host
	Host string `mapstructure:"host"`

	Project    Project      `mapstructure:"project"`
	Namespaces []*Namespace `mapstructure:"namespaces"`

	Server    Serve           `mapstructure:"serve"`
	Log       LogConfig       `mapstructure:"log"`
	Scheduler SchedulerConfig `mapstructure:"scheduler"`
	Telemetry TelemetryConfig `mapstructure:"telemetry"`

	namespaceNameToNamespace map[string]*Namespace
}

func (o *Optimus) GetNamespaceByName(name string) (*Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *Optimus) GetVersion() string { _ = "STUB: not implemented"; return "" }
