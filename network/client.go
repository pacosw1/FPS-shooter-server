package network

import (
	"encoding/json"
	"log"
	"sockets/entity"

	pb "sockets/protobuf"

	"github.com/golang/protobuf/proto"

	"sockets/events"
	"sockets/message"

	"github.com/gorilla/websocket"
)

func marshalMessage(message proto.Message) *[]byte {
	bytes, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}

	return &bytes
}

//Client t
type Client struct {
	ID   uint32
	Conn *websocket.Conn
}

//NewClient creates new client struct instance
func NewClient(ID uint32, c *websocket.Conn) *Client {
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
		e.FireInput(&m)

	}
}

func (c *Client) writeState(s *entity.Broadcast) {

	state := &pb.State{
		Players:     s.Players,
		Projectiles: s.Projectiles,
	}

	msg := marshalMessage(state)

	c.Conn.WriteMessage(websocket.BinaryMessage, *msg)

	// msg, err := proto.Marshall()

}
