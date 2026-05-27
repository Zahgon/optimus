package specio

import (
	"github.com/spf13/afero"

	"github.com/raystack/optimus/client/local"
	"github.com/raystack/optimus/client/local/model"
)

type jobSpecReadWriter struct {
	withParentReading bool

	referenceParentFileName string
	referenceSpecFileName   string
	referenceAssetDirName   string

	specFS afero.Fs
}

func NewJobSpecReadWriter(specFS afero.Fs, opts ...jobSpecReadWriterOpt) (local.SpecReadWriter[*model.JobSpec], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j jobSpecReadWriter) ReadAll(rootDirPath string) ([]*model.JobSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j jobSpecReadWriter) ReadByName(rootDirPath, name string) (*model.JobSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j jobSpecReadWriter) Write(dirPath string, spec *model.JobSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (j jobSpecReadWriter) writeJobSpecAsset(filePath, content string) error {
	_ = "STUB: not implemented"
	return nil
}

func (jobSpecReadWriter) mergeJobSpecWithParents(spec *model.JobSpec, specDirPath string, jobSpecParentsMappedByDirPath map[string]*model.JobSpec) {
	_ = "STUB: not implemented"
	return
}

func (j jobSpecReadWriter) readJobSpecParentsMappedByDirPath(rootDirPath string) (map[string]*model.JobSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j jobSpecReadWriter) readJobSpec(dirPath string) (*model.JobSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j jobSpecReadWriter) readJobSpecAssetsMappedByFileName(dirPath string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j jobSpecReadWriter) readJobSpecAssetFile(assetFilePath string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
