package img

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetFilenameParts(t *testing.T) {
	// Default case
	fileName, fileType := GetFilenameParts("test.jpg")
	if fileName != "test" {
		t.Log("File name was expected to be 'test' got", fileName)
		t.Fail()
	}
	if fileType != "jpg" {
		t.Log("File type was expected to be 'jpg' got", fileType)
		t.Fail()
	}

	// Case with multiple '.'
	fileName, fileType = GetFilenameParts("test.testing.jpg")
	if fileName != "test.testing" {
		t.Log("File name was expected to be 'test' got", fileName)
		t.Fail()
	}
	if fileType != "jpg" {
		t.Log("File type was expected to be 'jpg' got", fileType)
		t.Fail()
	}

	// Empty case
	fileName, fileType = GetFilenameParts("")
	t.Log("filename", fileName)
	if fileName != "" {
		t.Log("File name was expected to be 'test' got", fileName)
		t.Fail()
	}
	if fileType != "" {
		t.Log("File type was expected to be 'jpg' got", fileType)
		t.Fail()
	}
}

func TestCheckIfFileTypeIsSupported(t *testing.T) {
	fileTypes := strings.Split(SupportedFileTypes, "|")

	for _, x := range fileTypes {
		if CheckIfFileTypeIsSupported(fmt.Sprintf("%s.%s", "testFileName", x)) == false {
			t.Log("Got unsupported file type, expected supported file type")
			t.Fail()
		}
	}

	if CheckIfFileTypeIsSupported("testFileName.pdf") {
		t.Log("Got supported file type, expected unsupported file type")
		t.Fail()
	}
}
