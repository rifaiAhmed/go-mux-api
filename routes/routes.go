package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rifaiAhmed/go-rest-api-mux/controllers/gomuxController"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", gomuxController.Home)
	r.HandleFunc("/data", gomuxController.AllProducts)

	log.Fatal(http.ListenAndServe(":9080", r))
}
