package server

import (
	"sockets/events"
	"sockets/network"
	"sockets/simulation"
	"sockets/state"
)

//Server class that will hold everything together
type Server struct {
	GameState  *state.GameState
	EventQueue *events.EventQueue
	Simulation *simulation.PhysicsTicker
	Network    *network.Network
}

//New create a new server instance
func New() *Server {
	eventQ := events.NewEventQ()

	return &Server{
		EventQueue: eventQ,
	}

}

//Start start the server and all components
func (s *Server) Start() {
	println("Starting sever...")
	println()
	s.GameState = state.New(s.EventQueue)
	s.Simulation = simulation.NewPhysicsTicker(s.EventQueue)
	s.Network = network.New(s.EventQueue, s.GameState)

	// s.Simulation.fps = 60

	s.EventQueue.RegisterConnect(s.GameState)
	// s.EventQueue.RegisterProjectileReady(s.Simulation)

	//GameState listens for
	s.EventQueue.RegisterInput(s.GameState)
	s.EventQueue.RegisterTimeStep(s.GameState)

	// s.EventQueue.RegisterStartBroadcast(s.GameState)

	//Network listens for
	s.EventQueue.RegisterBroadcast(s.Network)
	s.EventQueue.RegisterDisconnect(s.Network)

	s.EventQueue.Start()
	s.GameState.Start()
	go s.Simulation.Run()
	s.Network.Start()

}

/*
1. Network sends client requests to event queue
2. EventQueue fires events, sends them to simulation
3. Simulation reads current Gamestate and updates physics
4. Simulation sends new GameState to eventQueue every 10 fps
5. Event Queue routes event back to Network.
6. Network broadcasts message to all clients
*/
