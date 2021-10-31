package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	img "./image"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/image/upload/{id}", img.HandleUpload).Methods("POST")
	myRouter.HandleFunc("/image/serve/{id}/{filename}", img.HandleServe).Methods("GET")
	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./demo/index.html") })
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	handleRequests()
}
