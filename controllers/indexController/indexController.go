package indexController

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shagtv/go-api/library/db"
	"github.com/shagtv/go-api/models"
	"github.com/shagtv/go-api/storage/userStorage"
	"math/rand"
	"net/http"
	"text/template"
)

func Index(res http.ResponseWriter, req *http.Request) {
	var data models.ClientNotes

	//db.Db.Offset(rand.Intn(13000)).Limit(1).Find(&data)
	db.Connection().Where("id = ?", rand.Intn(13000)).First(&data)

	dataString, err := json.Marshal(data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(dataString)
}

func Hello(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := params["name"]

	//req.URL.Query()

	//dataString := []byte("Hello " + name)

	t, _ := template.ParseFiles("templates/edit.html")
	t.Execute(res, name)
	//res.Header().Set("Content-Type", "text/html")
	//res.Write(dataString)
}

func Install(res http.ResponseWriter, req *http.Request) {
	userStorage.CreateTable()

	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte("Ok"))
}
