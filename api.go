package main

import (
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
)

/**
 * @name RegisterAndEnrollAPI
 * Gorilla Mux Api Function, Register and enrolls User
 * Returns JWT Token
 */
func RegisterAndEnrollAPI(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	json := simplejson.New()

	if username == "" || password == "" {

		json.Set("success", false)
		json.Set("message", "Please Supply username and password")

		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, err := RegisterAndEnrollUser(username, password)

		if err == nil {
			json.Set("success", true)
			json.Set("token", token)
			json.Set("message", "Succesfully registered and enrolled")

			w.WriteHeader(http.StatusOK)
		} else {
			json.Set("success", false)
			json.Set("message", err)

			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

/**
 * @name EnrollAPI
 * Gorilla Mux Api Function, Register and enrolls User
 * Returns JWT Token
 */
func EnrollAPI(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	json := simplejson.New()

	if username == "" || password == "" {

		json.Set("success", false)
		json.Set("message", "Please Supply username and password")

		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, err := EnrollUser(username, password)

		if err == nil {
			json.Set("success", true)
			json.Set("token", token)
			json.Set("message", "Succesfully registered and enrolled")

			w.WriteHeader(http.StatusOK)
		} else {
			json.Set("success", false)
			json.Set("message", err)

			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
