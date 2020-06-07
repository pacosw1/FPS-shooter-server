package network

import (
	"log"
	"sockets/entity"
	"sockets/utils"

	pb "sockets/protobuf"

	"github.com/golang/protobuf/proto"

	"sockets/events"
	"sockets/message"

	"github.com/gorilla/websocket"
)

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
		c.decodeClientMessage(msg, e)

	}
}

func (c *Client) decodeClientMessage(data []byte, e *events.EventQueue) {

	clientInput := &pb.Message{}

	if err := proto.Unmarshal(data, clientInput); err != nil {
		log.Fatalln("Failed to unmarshal UserInput:", err)
		return
	}

	switch t := clientInput.Payload.(type) {
	case *pb.Message_PlayerInput:
		input := clientInput.GetPlayerInput()
		e.FireInput(&message.NetworkInput{
			ID:         c.ID,
			IsShooting: input.IsShooting,
			Direction:  message.ProtoPoint(input.Direction),
			SequenceID: input.SequenceID,
			Rotation:   message.ProtoVector(input.Rotation),
		})
	default:
		log.Println("Unknown message type", t)
	}

}

func (c *Client) writeState(s *entity.Broadcast) {

	state := &pb.State{
		Players:     s.Players,
		Projectiles: s.Projectiles,
	}

	msg := utils.MarshalMessage(state)

	c.Conn.WriteMessage(websocket.BinaryMessage, *msg)

	// msg, err := proto.Marshall()

}
