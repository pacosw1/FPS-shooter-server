package state

import (
	"sockets/entity"
	"sockets/message"
)

//New game state constructor
func New() *GameState {
	return &GameState{
		Players:     make(map[int]*entity.Player),
		Projectiles: make(map[int16]*entity.Projectile),
	}
}

func (g *GameState) StateBroadcast() {

	for len(g.Players) > 0 {
		
	}
}

//GameState Whole game state
type GameState struct {
	requests    chan message.UserInput
	Players     map[int]*entity.Player
	Projectiles map[int16]*entity.Projectile
}

//HandleInput request
func (g *GameState) HandleInput(m *message.UserInput) {
	g.Players[m.ID].UpdatePlayer(m)
}

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
	println(len(g.Players))

}

//HandleDisconnect disconnect player
func (g *GameState) HandleDisconnect(m *message.Disconnect) {
	g.RemovePlayer(m)
}

func (g *GameState) updatePlayer(m *message.UserInput) {
	player, exists := g.Players[m.ID]

	if exists {
		player.UpdatePlayer(m)
	}

}
