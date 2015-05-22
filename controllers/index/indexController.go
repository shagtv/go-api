package indexController

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shagtv/go-api/library/db"
	"github.com/shagtv/go-api/models"
	"math/rand"
)

func Index(res http.ResponseWriter, req *http.Request) {
	var data models.ClientNotes

	//db.Db.Offset(rand.Intn(13000)).Limit(1).Find(&data)
	db.Db.Where("id = ?", rand.Intn(13000)).First(&data)

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

	dataString := []byte("Hello " + name)

	res.Header().Set("Content-Type", "text/html")
	res.Write(dataString)
}