package network

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//CreateNetwork Initialize Network structure
func CreateNetwork() *Network {
	return &Network{
		Clients: make(map[string]*websocket.Conn),
		Total:   0,
	}
}

//Network that will hold client data
type Network struct {
	Clients map[string]*websocket.Conn
	Total   int
}

//upgrades initial http request to websocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Request struct {
	Type     int
	ClientID string
	Data     int
}

//socket hanfles socket connection and data stream
func Socket(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	// print(conn)

}
