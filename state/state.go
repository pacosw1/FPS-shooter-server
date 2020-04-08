package state

import (
	"sockets/entity"
	"sockets/events"
	"sockets/message"
	"time"
)

//New game state constructor
func New(e *events.EventQueue) *GameState {
	return &GameState{
		Players:     make(map[int]*entity.Player),
		Projectiles: make(map[int16]*entity.Projectile),
		EventQueue:  e,
	}
}

//Start the broadcast timer
func (g *GameState) Start() {
	println("State Updater Online")
	seconds := time.Duration(1000 / 10)
	ticker := time.Tick(seconds * time.Millisecond)
	go g.broadcastState(ticker)

}
func (g *GameState) broadcastState(t <-chan time.Time) {

	for {
		select {
		case <-t:
			g.EventQueue.FireGameState(message.SendState())
		}

	}
}

//GameState Whole game state
type GameState struct {
	requests    chan message.UserInput
	Players     map[int]*entity.Player
	Projectiles map[int16]*entity.Projectile
	EventQueue  *events.EventQueue
}

//HandleInput request
func (g *GameState) HandleInput(m *message.NetworkInput) {
	g.Players[m.ID].UpdatePlayer(m)
}

//RemovePlayer removes player
func (g *GameState) RemovePlayer(m *message.Disconnect) {
	delete(g.Players, m.ClientID)
	println(len(g.Players))

}

//AddPlayer 1
func (g *GameState) AddPlayer(m *message.Connect) {
	_, exists := g.Players[m.ClientID]
	if !exists {
		g.Players[m.ClientID] = entity.NewPlayer(m.ClientID)
	}
}

//HandleConnect add player on connect request
func (g *GameState) HandleConnect(m *message.Connect) {

	g.AddPlayer(m)
	println("New player connected, total: ", len(g.Players))

}

//HandleDisconnect disconnect player
func (g *GameState) HandleDisconnect(m *message.Disconnect) {
	g.RemovePlayer(m)
}

func (g *GameState) updatePlayer(m *message.NetworkInput) {
	player, exists := g.Players[m.ID]

	if exists {
		player.UpdatePlayer(m)
	}

}
