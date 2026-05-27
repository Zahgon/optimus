package survey

import (
	"github.com/spf13/afero"
)

const (
	answerYes = "Yes"
	answerNo  = "No"
)

// AskWorkingDirectory asks and returns the directory where the new spec folder should be created
func AskWorkingDirectory(specFS afero.Fs, rootDirPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AskDirectoryName asks and returns the directory name of the new spec folder
func AskDirectoryName(root string) (string, error) { _ = "STUB: not implemented"; return "", nil }
