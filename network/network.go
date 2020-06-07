package network

import (
	"net/http"
	"sockets/entity"
	"sockets/events"
	"sockets/message"
	"sockets/state"
	"sockets/validate"

	"github.com/gorilla/websocket"
)

//Network that will hold client data
type Network struct {
	Clients   map[uint32]*Client
	EventQ    *events.EventQueue
	GameState *state.GameState
}

//AddClient adds client to network
func (n *Network) AddClient(c *websocket.Conn) *Client {
	newID := PlayerID(100, n)
	client := NewClient(newID, c)
	n.Clients[newID] = client
	return n.Clients[newID]
}

//New Initialize Network structure
func New(e *events.EventQueue, g *state.GameState) *Network {
	return &Network{
		Clients:   make(map[uint32]*Client),
		EventQ:    e,
		GameState: g,
	}
}

//PlayerID test
func PlayerID(size int, n *Network) uint32 {
	uniqueID := validate.GenerateID(size)
	if _, ok := n.Clients[uniqueID]; ok {
		uniqueID = PlayerID(size, n)
	}
	return uniqueID
}

//HandleDisconnect closes connection and diconnects player
func (n *Network) HandleDisconnect(m *message.Disconnect) {
	n.Clients[m.ClientID].Conn.Close()
	delete(n.GameState.Players, m.ClientID)
}

//HandleStateBroadcast t
func (n *Network) HandleStateBroadcast(m *entity.Broadcast) {
	n.broadcastState(m)
}

func (n *Network) broadcastState(s *entity.Broadcast) {
	for _, client := range n.Clients {
		client.writeState(s)
	}
}

//Start starts new network
func (n *Network) Start() {
	println("Network Online")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/socket", n.Socket)
	http.ListenAndServe(":5000", nil)

}

//RemoveClient removes client from network
func (n *Network) RemoveClient(ID uint32) {
	delete(n.Clients, ID)
}
