package network

import (
	"encoding/json"
	"log"
	"net/http"
	"sockets/message"
	"sockets/state"

	"github.com/gorilla/websocket"
)

//upgrades initial http request to websocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//AddClient adds client to network
func (n *Network) AddClient(c *websocket.Conn) *Client {
	newID := PlayerID(100, n)
	client := NewClient(newID, c)
	n.Clients[newID] = client
	return n.Clients[newID]
}

//Listen listens for client input
func (c *Client) Listen() {
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		var m message.NetworkMessage
		mErr := json.Unmarshal(msg, &m)
		if mErr != nil {
			return
		}
	}
}

func (c *Client) writeState(s *state.GameState) {
	c.Conn.WriteJSON(s.Players)
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
	er := client.Conn.WriteJSON(message.ConnectMessage("pacosw1", 1))
	if er != nil {
		println("error sending id")
	}
	n.EventQ.FireConnect(message.ConnectMessage("pacosw1", client.ID))

	defer conn.Close()

	client.Listen()

	// err = conn.WriteMessage(mt, message)
	// if err != nil {
	// 	log.Println("write:", err)
	// 	break
	// }

	// print(conn)

}
