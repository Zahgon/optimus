package specio

import (
	"github.com/spf13/afero"

	"github.com/raystack/optimus/client/local"
	"github.com/raystack/optimus/client/local/model"
)

type resourceSpecReadWriter struct {
	referenceSpecFileName string
	specFS                afero.Fs
}

func NewResourceSpecReadWriter(specFS afero.Fs) (local.SpecReadWriter[*model.ResourceSpec], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r resourceSpecReadWriter) ReadAll(rootDirPath string) ([]*model.ResourceSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: in the future, we should make it so that we can identify the resource exist or not based on the file path
func (r resourceSpecReadWriter) ReadByName(rootDirPath, name string) (*model.ResourceSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r resourceSpecReadWriter) Write(dirPath string, spec *model.ResourceSpec) error {
	_ = "STUB: not implemented"
	return nil
}
