package main

import (
    "net/http"
    "api/controllers/indexController"
    "api/controllers/brandController"
    "github.com/gorilla/mux"
    "api/library/db"
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