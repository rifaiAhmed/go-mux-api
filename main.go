package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rifaiAhmed/go-rest-api-mux/controllers/gomuxController"
)

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", gomuxController.Home)
	r.HandleFunc("/product", gomuxController.AllProducts)

	log.Fatal(http.ListenAndServe(":9080", r))
}
func main() {
	handleRequest()
}
