package extension

import (
	"io"

	"github.com/spf13/afero"

	"github.com/raystack/optimus/client/extension/model"
)

// AssetOperatorFS is file system that will be used by operator.
// It can be changed before calling any operation.
// But, make sure to change it back after the operation is done
// to its default value to avoid unexpected behaviour.
var AssetOperatorFS = afero.NewOsFs()

type defaultAssetOperator struct {
	localDirPath string

	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

// NewDefaultAssetOperator initializes default asset operator
func NewDefaultAssetOperator(stdin io.Reader, stdout, stderr io.Writer) model.AssetOperator {
	_ = "STUB: not implemented"
	return *new(model.AssetOperator)
}

func (d *defaultAssetOperator) Prepare(localDirPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultAssetOperator) Install(asset []byte, tagName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultAssetOperator) Uninstall(tagNames ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultAssetOperator) Run(tagName string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}
