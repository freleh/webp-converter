package img

import (
	"os"
	"strings"
)

const (
	supportedFileTypes = "jpeg|jpg|png|webp"
)

// Check if the directory dir exists.
// If not, it gets created.
func CheckAndCreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0777); err != nil {
			return err
		}
	}
	return nil
}

// Returns true if file type is supported, otherwise returns false
func CheckIfFileTypeIsSupported(filename string) bool {
	filenameParts := strings.Split(filename, ".")
	fileType := filenameParts[len(filenameParts)-1]

	fileTypes := strings.Split(supportedFileTypes, "|")
	for _, x := range fileTypes {
		if x == fileType {
			return true
		}
	}

	return false
}
