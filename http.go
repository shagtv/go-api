package main

import (
	"github.com/gorilla/mux"
	"github.com/shagtv/go-api/controllers/brandController"
	"github.com/shagtv/go-api/controllers/indexController"
	"github.com/shagtv/go-api/controllers/userController"
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
	rtr.HandleFunc("/hello/{name}", indexController.Hello).Methods("GET")
	rtr.HandleFunc("/install", indexController.Install).Methods("GET")

	rtr.HandleFunc("/user/login", userController.Login)
	rtr.HandleFunc("/user/logout", userController.Logout)
	rtr.HandleFunc("/user/register", userController.Register)
	rtr.HandleFunc("/user/unregister", userController.Unregister)

	rtr.HandleFunc("/brand/info", brandController.Info).Methods("GET")
	rtr.HandleFunc("/brand/list", brandController.List).Methods("GET")
	rtr.HandleFunc("/brand/count", brandController.Count).Methods("GET")

	rtr.HandleFunc("/{name}", indexController.Hello).Methods("GET")
}
