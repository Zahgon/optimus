package plugin

const (
	TypeTask Type = "task"
	TypeHook Type = "hook"

	HookTypePre  HookType = "pre"
	HookTypePost HookType = "post"
	HookTypeFail HookType = "fail"

	ModTypeCLI                Mod = "cli"
	ModTypeDependencyResolver Mod = "dependencyresolver"

	DestinationURNFormat = "%s://%s"
)

type Type string

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

type HookType string

func (ht HookType) String() string { _ = "STUB: not implemented"; return "" }

type Mod string

func (m Mod) String() string { _ = "STUB: not implemented"; return "" }

type Entrypoint struct {
	Shell  string
	Script string
}

type Info struct {
	// Name should as simple as possible with no special characters
	// should start with a character, better if all lowercase
	Name        string
	Description string
	PluginType  Type  `yaml:",omitempty"`
	PluginMods  []Mod `yaml:",omitempty"`

	PluginVersion string   `yaml:",omitempty"`
	APIVersion    []string `yaml:",omitempty"`

	// Image is the full path to docker container that will be scheduled for execution
	Image string

	// Entrypoint command which will be used to execute the plugin
	Entrypoint Entrypoint

	// DependsOn returns list of hooks this should be executed after
	DependsOn []string `yaml:",omitempty"`

	// PluginType provides the place of execution, could be before the transformation
	// after the transformation, etc
	HookType HookType `yaml:",omitempty"`
}

func (info *Info) Validate() error { _ = "STUB: not implemented"; return nil }

// image is a required field

// version is a required field

// entrypoint is a required field

type YamlMod interface {
	PluginInfo() *Info
	CommandLineMod
}

type Config struct {
	Name  string
	Value string
}

type Configs []Config

func (c Configs) Get(name string) (Config, bool) {
	_ = "STUB: not implemented"
	return *new(Config), false
}

func ConfigsFromMap(configMap map[string]string) Configs {
	_ = "STUB: not implemented"
	return *new(Configs)
}

type Asset struct {
	Name  string
	Value string
}

type Assets []Asset

func AssetsFromMap(assetsMap map[string]string) Assets {
	_ = "STUB: not implemented"
	return *new(Assets)
}

func (a Assets) Get(name string) (Asset, bool) {
	_ = "STUB: not implemented"
	return *new(Asset), false
}

func (a Assets) ToMap() map[string]string { _ = "STUB: not implemented"; return nil }

// Plugin is an extensible module implemented outside the core optimus boundaries
type Plugin struct {
	// Mods apply multiple modifications to existing registered plugins which
	// can be used in different circumstances
	DependencyMod DependencyResolverMod
	YamlMod       YamlMod
}

func (p *Plugin) IsYamlPlugin() bool { _ = "STUB: not implemented"; return false }

func (p *Plugin) GetSurveyMod() CommandLineMod {
	_ = "STUB: not implemented"
	return *new(CommandLineMod)
}

func (p *Plugin) Info() *Info { _ = "STUB: not implemented"; return nil }
