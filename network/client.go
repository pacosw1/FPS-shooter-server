package network

import (
	"encoding/json"
	"log"
	"sockets/events"
	"sockets/message"
	"sockets/state"

	"github.com/gorilla/websocket"
)

//Client t
type Client struct {
	ID   int
	Conn *websocket.Conn
}

//NewClient creates new client struct instance
func NewClient(ID int, c *websocket.Conn) *Client {
	return &Client{
		ID:   ID,
		Conn: c,
	}
}

//Listen listens
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