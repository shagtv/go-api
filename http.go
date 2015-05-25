package main

import (
	"github.com/gorilla/mux"
	"github.com/shagtv/go-api/controllers/brand"
	"net/http"
)

func main() {
	rtr := mux.NewRouter()
	fillRouter(rtr)

	http.Handle("/", rtr)
	http.ListenAndServe(":9000", nil)
}

func fillRouter(rtr *mux.Router) {
	rtr.HandleFunc("/", brandController.Count).Methods("GET")

	rtr.HandleFunc("/brand/info", brandController.Info).Methods("GET")
	rtr.HandleFunc("/brand/list", brandController.List).Methods("GET")
	rtr.HandleFunc("/brand/count", brandController.Count).Methods("GET")
}
