package img

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	dir := fmt.Sprintf("%s%s", "./images/", id)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.Mkdir(dir, 0666)
	}

}
