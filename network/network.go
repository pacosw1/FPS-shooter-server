package network

import (
	"net/http"
	"sockets/events"
	"sockets/message"
	"sockets/state"
)

//New Initialize Network structure
func New(e *events.EventQueue, g *state.GameState) *Network {
	return &Network{
		Clients:   make(map[int]*Client),
		EventQ:    e,
		GameState: g,
	}
}

//HandleStateBroadcast t
func (n *Network) HandleStateBroadcast(m *message.StateMessage) {
	println("broadcasting state to clients")
	n.broadcastState()
}

func (n *Network) broadcastState() {
	for _, client := range n.Clients {
		client.writeState(n.GameState)
	}
}

//Start starts new network
func (n *Network) Start() {

	http.HandleFunc("/socket", n.Socket)
	http.ListenAndServe(":8080", nil)

}

//Network that will hold client data
type Network struct {
	Clients   map[int]*Client
	EventQ    *events.EventQueue
	GameState *state.GameState
}

//RemoveClient removes client from network
func (n *Network) RemoveClient(ID int) {
	delete(n.Clients, ID)
}
