package extension

import (
	"github.com/spf13/afero"
)

// CleanExtensionFS is file system that will be used when cleaning extension.
// It can be changed before calling the clean operation.
// But, make sure to change it back after the operation is done
// to its default value to avoid unexpected behaviour.
var CleanExtensionFS = afero.NewOsFs()

// Clean cleans all extensions from local, including its manifest
func Clean(verbose bool) error { _ = "STUB: not implemented"; return nil }
