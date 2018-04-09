package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	createLogger()
	initKeys()
	bcNetworksetup()
	enrollAdmin()

	router := mux.NewRouter()
	router.HandleFunc("/registerandenroll", RegisterAndEnrollAPI).Methods("POST")
	router.HandleFunc("/enroll", EnrollAPI).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
