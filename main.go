package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorld()).Methods("GET")
	router.HandleFunc("/goodbye", goodbyeMoonMen()).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func helloWorld() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	}
}

func goodbyeMoonMen() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Goodbye Moon Men"))
	}
}
