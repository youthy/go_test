package util

import (
	"log"
	"net/http"
)

// CheckErr log err and panic
func CheckErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

// ServerErr response http code 500
func ServerErr(w http.ResponseWriter) {
	if p := recover(); p != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - something oops happend on server, sorry for that"))
	}
}

// BadRequest response http code 400
func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 - Bad Request"))
}
