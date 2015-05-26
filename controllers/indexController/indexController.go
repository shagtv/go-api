package indexController

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/shagtv/go-api/library/db"
	"github.com/shagtv/go-api/models"
	"github.com/shagtv/go-api/storage/userStorage"
	"log"
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

var Connections map[*websocket.Conn]bool

func sendAll(msg []byte) {
	for conn := range Connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(Connections, conn)
			conn.Close()
		}
	}
}

func Ws(res http.ResponseWriter, req *http.Request) {
	// Taken from gorilla's website
	conn, err := websocket.Upgrade(res, req, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(res, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	log.Println("Succesfully upgraded connection")
	Connections[conn] = true

	for {
		// Blocks until a message is read
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(Connections, conn)
			conn.Close()
			return
		}
		log.Println(string(msg))
		sendAll(msg)
	}
}

func Install(res http.ResponseWriter, req *http.Request) {
	userStorage.CreateTable()

	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte("Ok"))
}
