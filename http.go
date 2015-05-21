package main

import (
    "net/http"
    "github.com/shagtv/go-api/controllers/indexController"
    "github.com/shagtv/go-api/controllers/brandController"
    "github.com/gorilla/mux"
    "github.com/shagtv/go-api/library/db"
)

func main() {
    rtr := mux.NewRouter()
    fillRouter(rtr);

    db.Connect()

    http.Handle("/", rtr)
    http.ListenAndServe(":9000", nil)
}

func fillRouter(rtr *mux.Router) {
    rtr.HandleFunc("/", indexController.Index).Methods("GET")
    rtr.HandleFunc("/hello/{name:[a-z]+}", indexController.Hello).Methods("GET")

    rtr.HandleFunc("/brand/info", brandController.Info).Methods("GET")
}