package server

import (
	"sockets/events"
	"sockets/network"
	"sockets/state"
)

//Server class that will hold everything together
type Server struct {
	GameState  *state.GameState
	EventQueue *events.EventQueue
	Simulation *simulation.Engine
	Network    *network.Network
}

//Start Init server components
func Start(s *Server) {
}

/*
1. Network sends client requests to event queue
2. EventQueue fires events, sends them to simulation
3. Simulation reads current Gamestate and updates physics
4. Simulation sends new GameState to eventQueue every 10 fps
5. Event Queue routes event back to Network.
6. Network broadcasts message to all clients
*/
