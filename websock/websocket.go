package websock

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// We set our Read and Write buffer sizes for our uprader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Upgrade function will take an incoming request and upgrade it to a websocket connection
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// this line allows other origin hosts to connect to our
	// websocket server
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// create our websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	// return our new websocket connection
	return ws, nil
}
