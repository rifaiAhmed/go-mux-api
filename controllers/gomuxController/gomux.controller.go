package gomuxController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rifaiAhmed/go-rest-api-mux/services/gomuxService"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "selamat datang")
}
func AllProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(gomuxService.GetData())
}
