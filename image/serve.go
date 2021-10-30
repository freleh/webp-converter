package img

import (
	"image"
	"log"
	"net/http"
	"os"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"

	_ "image/jpeg"
	_ "image/png"
)

func HandleServe(w http.ResponseWriter, r *http.Request) {
	reader, err := os.Open("images/tangerine-flower.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("images/tangerine-flower.webp")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}

	if err := webp.Encode(f, m, options); err != nil {
		log.Fatalln(err)
	}
}
