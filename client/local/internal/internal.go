package internal

import (
	"io/fs"

	"github.com/spf13/afero"

	"github.com/raystack/optimus/client/local"
)

func DiscoverSpecDirPaths(specFS afero.Fs, rootSpecDir, referenceFileName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DiscoverFilePaths(fileFS afero.Fs, rootDir string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func discoverPathsUsingSelector(specFS afero.Fs, rootSpecDir string, selectPath func(path string, info fs.FileInfo) (string, bool)) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WriteSpec[S local.ValidSpec](specFS afero.Fs, filePath string, spec S) error {
	_ = "STUB: not implemented"
	return nil
}

func ReadSpec[S local.ValidSpec](specFS afero.Fs, filePath string) (S, error) {
	_ = "STUB: not implemented"
	return *new(S), nil
}

func GetFirstSpecByFilter[S local.ValidSpec](specs []S, filter func(S) bool) S {
	_ = "STUB: not implemented"
	return *new(S)
}
