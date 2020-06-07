package network

import (
	"log"
	"net/http"
	"sockets/message"
	pb "sockets/protobuf"
	"sockets/utils"

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

	response := &pb.ConnectResponse{
		Success:  false,
		Username: "test",
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {

		conn.WriteMessage(websocket.BinaryMessage, *utils.MarshalMessage((response)))
		log.Fatal(err)
		return
	}

	client := n.AddClient(conn)
	//successfull connection
	response.Id = client.ID
	response.Success = true

	msg := utils.MarshalMessage(response)

	err = conn.WriteMessage(websocket.BinaryMessage, *msg)
	n.EventQ.FireConnect(message.ConnectMessage("test", client.ID))

	defer conn.Close()

	client.Listen(n.EventQ)
}
