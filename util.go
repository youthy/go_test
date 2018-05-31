package main

import (
	"log"
	"net/http"
)

// if err then panic
func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func serverErr(w http.ResponseWriter) {
	if p := recover(); p != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - something oops happend on server, sorry for that"))
	}
}

func badRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 - Bad Request"))
}
