package img

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	// Read ID
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if directory exists, if not create one
	dir := fmt.Sprintf("%s%s", "./images/", id)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0777); err != nil {
			log.Printf("Could not create directory at: HandleUpload")
			w.WriteHeader(500)
		}
	}

	// Save file to directory

	//10 << 20 specifies a maximum upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Printf("No file with key: image at: HandleUpload")
		w.WriteHeader(422)
	}
	defer file.Close()

	// Create target file
	targetFile, err := os.Create(fmt.Sprintf("%s/%s", dir, handler.Filename))
	if err != nil {
		log.Printf("Could not create file at: HandleUpload with: %s/%s because: %s", dir, handler.Filename, err)
		w.WriteHeader(500)
	}
	defer targetFile.Close()

	// read all of the contents of our uploaded file into a byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Could not read bytes at: HandleUpload")
		w.WriteHeader(500)
	}

	targetFile.Write(fileBytes)
}
