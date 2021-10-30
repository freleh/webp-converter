package img

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	// Read ID
	vars := mux.Vars(r)
	id := vars["id"]

	//10 << 20 specifies a maximum upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check if the file type is supported
	if !checkIfFileTypeIsSupported(handler.Filename) {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// File directory
	dir := fmt.Sprintf("%s%s", "./images/", id)
	if err := checkAndCreateDir(dir); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create target file
	localFilename := fmt.Sprintf("%s/%s", dir, handler.Filename)
	targetFile, err := os.Create(localFilename)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	// read all of the contents of our uploaded file into a byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	targetFile.Write(fileBytes)
}

// Check if the directory dir exists.
// If not, it gets created.
func checkAndCreateDir(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0777); err != nil {
			return err
		}
	}
	return nil
}

// Returns true if file type is supported, otherwise returns false
func checkIfFileTypeIsSupported(filename string) bool {
	supportedFileTypes := [4]string{"jpeg", "jpg", "png", "webp"}

	filenameParts := strings.Split(filename, ".")
	fileType := filenameParts[len(filenameParts)-1]

	for _, x := range supportedFileTypes {
		if x == fileType {
			return true
		}
	}

	return false
}
