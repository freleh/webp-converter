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
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	handleRequests()
}
