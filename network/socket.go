package network

import (
	"encoding/json"
	"log"
	"net/http"
	"sockets/events"
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
func (c *Client) Listen(e *events.EventQueue) {
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			e.FireDisconnect(message.DisconnectMessage(c.ID))
			break
		}
		var m message.NetworkInput
		mErr := json.Unmarshal(msg, &m)
		if mErr != nil {
			e.FireDisconnect(message.DisconnectMessage(c.ID))
		}
		e.FireInput(message.SendInput(&m))

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
	connectMsg := message.ConnectMessage("pacosw1", client.ID)
	err = conn.WriteJSON(connectMsg)
	n.EventQ.FireConnect(connectMsg)

	defer conn.Close()

	client.Listen(n.EventQ)

	// err = conn.WriteMessage(mt, message)
	// if err != nil {
	// 	log.Println("write:", err)
	// 	break
	// }

	// print(conn)

}
