package secret

import (
	"time"

	"github.com/spf13/cobra"
)

const (
	secretTimeout = time.Minute * 2

	// TODO: get rid of system defined secrets
	systemDefinedSecretPrefix = "_OPTIMUS_"
)

// NewSecretCommand initializes command for secret
func NewSecretCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getSecretName(args []string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getSecretValue(args []string, filePath string, encoded bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//nolint: gomnd

func validateProperlyEncoded(secretValue string) error { _ = "STUB: not implemented"; return nil }
