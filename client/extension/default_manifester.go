package extension

import (
	"github.com/spf13/afero"

	"github.com/raystack/optimus/client/extension/model"
)

const manifestFileName = "manifest.yaml"

// ManifesterFS is file system that will be used for manifester operations.
// It can be changed before calling any manifester operation.
// But, make sure to change it back after the operation is done
// to its default value to avoid unexpected behaviour.
var ManifesterFS = afero.NewOsFs()

type defaultManifester struct{}

// NewDefaultManifester initializes default manifester
func NewDefaultManifester() model.Manifester {
	_ = "STUB: not implemented"
	return *new(model.Manifester)
}

// Load loads manifest from local machine
func (d *defaultManifester) Load(dirPath string) (*model.Manifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*defaultManifester) enrichManifest(manifest *model.Manifest) {
	_ = "STUB: not implemented"
	return
}

// Flush flushes manifest into a file in local machine
func (*defaultManifester) Flush(manifest *model.Manifest, dirPath string) error {
	_ = "STUB: not implemented"
	return nil
}
