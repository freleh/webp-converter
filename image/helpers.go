package img

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	SupportedFileTypes = "jpeg|jpg|png|webp"
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
	_, fileType := GetFilenameParts(filename)

	fileTypes := strings.Split(SupportedFileTypes, "|")
	for _, x := range fileTypes {
		if x == fileType {
			return true
		}
	}

	return false
}

// Split filename into name and type
func GetFilenameParts(filename string) (name string, fileType string) {
	filenameParts := strings.Split(filename, ".")
	return strings.Join(filenameParts[:len(filenameParts)-1], "."), filenameParts[len(filenameParts)-1]
}

// Check if file exists
func FileExists(name string) (string, error) {
	fileTypes := strings.Split(SupportedFileTypes, "|")

	for _, x := range fileTypes {
		searchFile := fmt.Sprintf("%s.%s", name, x)
		_, err := os.Stat(searchFile)
		if !errors.Is(err, os.ErrNotExist) {
			return searchFile, nil
		}
	}

	return "", fmt.Errorf("file %s not found", name)
}
