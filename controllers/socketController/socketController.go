package socketController

import (
	"github.com/gorilla/websocket"
	"github.com/shagtv/go-api/storage/userStorage"
	"log"
	"net/http"
)

var Connections map[*websocket.Conn]string

func sendAll(msg []byte, addr string) {
	for conn, address := range Connections {
		if address == addr {
			continue
		}
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(Connections, conn)
			conn.Close()
		}
	}
}

func Ws(res http.ResponseWriter, req *http.Request) {

	if Connections == nil {
		Connections = make(map[*websocket.Conn]string)
	}

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
	Connections[conn] = conn.RemoteAddr().String()

	log.Println(conn.RemoteAddr())
	if err := conn.WriteMessage(websocket.TextMessage, []byte(conn.RemoteAddr().String())); err != nil {
		delete(Connections, conn)
		conn.Close()
	}

	for {
		// Blocks until a message is read
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(Connections, conn)
			conn.Close()
			return
		}
		log.Println(string(msg))
		sendAll(msg, conn.RemoteAddr().String())
	}
}

func Install(res http.ResponseWriter, req *http.Request) {
	userStorage.CreateTable()

	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte("Ok"))
}
