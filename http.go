package main

import (
	"github.com/gorilla/mux"
	"github.com/shagtv/go-api/controllers/brandController"
	"github.com/shagtv/go-api/controllers/indexController"
	"github.com/shagtv/go-api/controllers/socketController"
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
	rtr.HandleFunc("/app/", brandController.Count).Methods("GET")
	rtr.HandleFunc("/app/hello/{name}", indexController.Hello).Methods("GET")
	rtr.HandleFunc("/app/install", indexController.Install).Methods("GET")

	rtr.HandleFunc("/ws", socketController.Ws)

	rtr.HandleFunc("/app/user/login", userController.Login)
	rtr.HandleFunc("/app/user/logout", userController.Logout)
	rtr.HandleFunc("/app/user/register", userController.Register)
	rtr.HandleFunc("/app/user/unregister", userController.Unregister)

	rtr.HandleFunc("/app/brand/info", brandController.Info).Methods("GET")
	rtr.HandleFunc("/app/brand/list", brandController.List).Methods("GET")
	rtr.HandleFunc("/app/brand/count", brandController.Count).Methods("GET")

	rtr.HandleFunc("/app/{name}", indexController.Hello).Methods("GET")
}
