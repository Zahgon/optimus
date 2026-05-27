package plugin

import (
	"net/url"

	"github.com/raystack/salt/log"
	"github.com/spf13/cobra"

	"github.com/raystack/optimus/config"
)

type syncCommand struct {
	configFilePath string
	clientConfig   *config.ClientConfig
	logger         log.Logger
}

func NewSyncCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (s *syncCommand) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func getPluginDownloadURL(host string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *syncCommand) downloadArchiveFromServer() error { _ = "STUB: not implemented"; return nil }

func (s *syncCommand) RunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}
