package config

import (
	"github.com/spf13/afero"
)

const (
	DefaultFilename       = "optimus.yaml"
	DefaultConfigFilename = "config.yaml" // default file name for server config
	DefaultFileExtension  = "yaml"
	DefaultEnvPrefix      = "OPTIMUS"
	EmptyPath             = ""
)

var FS = afero.NewReadOnlyFs(afero.NewOsFs())

// LoadClientConfig load the project specific config from these locations:
// 1. filepath. ./optimus <client_command> -c "path/to/config/optimus.yaml"
// 2. current dir. Optimus will look at current directory if there's optimus.yaml there, use it
func LoadClientConfig(filePath string) (*ClientConfig, error) {
	_ = "STUB: not implemented"
	return nil,

		// getViperWithDefault + SetFs
		nil
}

// load opt from filepath if exist

// if filepath not valid, returns err

// load opt from current directory

// load the config

// LoadServerConfig load the server specific config from these locations:
// 1. filepath. ./optimus <server_command> -c "path/to/config.yaml"
// 2. env var. eg. OPTIMUS_SERVE_PORT, etc
// 3. executable binary location
func LoadServerConfig(filePath string) (*ServerConfig, error) {
	_ = "STUB: not implemented"
	return nil,

		// getViperWithDefault + SetFs
		nil
}

// load opt from filepath if exist

// if filepath not valid, returns err

// load opt from exec

// load opt from env var

// load the config

func validateFilepath(fs afero.Fs, fpath string) error { _ = "STUB: not implemented"; return nil }
