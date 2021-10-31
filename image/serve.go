package img

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kolesa-team/go-webp/decoder"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"

	jpeg "image/jpeg"
	png "image/png"
)

func HandleServe(w http.ResponseWriter, r *http.Request) {
	// Read ID
	vars := mux.Vars(r)
	id := vars["id"]
	name, targetType := GetFilenameParts(vars["filename"])

	fullFilePath, err := FileExists(fmt.Sprintf("./images/%s/%s", id, name))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	reader, err := os.Open(fullFilePath)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer reader.Close()

	var m image.Image
	_, sourceType := GetFilenameParts(fullFilePath)

	if sourceType == targetType {
		http.ServeFile(w, r, fullFilePath)
		return
	}

	if sourceType == "webp" {
		m, err = webp.Decode(reader, &decoder.Options{})
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		m, _, err = image.Decode(reader)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	f, err := os.CreateTemp("", fmt.Sprintf("*%s.%s", name, targetType))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer os.Remove(f.Name())

	if targetType == "webp" {
		options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := webp.Encode(f, m, options); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if targetType == "jpeg" || targetType == "jpg" {
		if err := jpeg.Encode(f, m, &jpeg.Options{}); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if targetType == "png" {
		if err := png.Encode(f, m); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	http.ServeFile(w, r, f.Name())
}
