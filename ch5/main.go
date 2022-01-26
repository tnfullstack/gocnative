package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, net/http!\n"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HelloHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
