package network

import (
	"log"
	"net/http"
	"sockets/message"

	"github.com/gorilla/websocket"
)

//upgrades initial http request to websocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Socket hanfles socket connection and data stream
func (n *Network) Socket(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	client := n.AddClient(conn)
	connectMsg := message.ConnectMessage("pacosw1", client.ID)
	err = conn.WriteJSON(connectMsg)
	n.EventQ.FireConnect(connectMsg)

	defer conn.Close()

	client.Listen(n.EventQ)
}
