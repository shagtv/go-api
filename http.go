package main

import (
	"github.com/gorilla/mux"
	"github.com/shagtv/go-api/controllers/brand"
	"github.com/shagtv/go-api/controllers/index"
	"github.com/shagtv/go-api/library/db"
	"net/http"
)

func main() {
	rtr := mux.NewRouter()
	fillRouter(rtr)

	db.Connect()

	http.Handle("/", rtr)
	http.ListenAndServe(":9000", nil)
}

func fillRouter(rtr *mux.Router) {
	rtr.HandleFunc("/", indexController.Index).Methods("GET")
	rtr.HandleFunc("/hello/{name:[a-z]+}", indexController.Hello).Methods("GET")

	rtr.HandleFunc("/brand/info", brandController.Info).Methods("GET")
	rtr.HandleFunc("/brand/list", brandController.List).Methods("GET")
	rtr.HandleFunc("/brand/count", brandController.Count).Methods("GET")
}
