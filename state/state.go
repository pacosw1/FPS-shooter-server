package state

import (
	"sockets/entity"
	"sockets/events"
	"sockets/message"
	"sockets/types"
	"sockets/validate"
	"time"
)

//New game state constructor
func New(e *events.EventQueue) *GameState {
	return &GameState{
		Players:     make(map[uint32]*entity.Player),
		Projectiles: make(map[uint32]*entity.Projectile),
		// Zombies:     make(map[int]*entity.Zombie),
		EventQueue: e,
	}
}

//Start the broadcast timer
func (g *GameState) Start() {
	println("State Updater Online")
	// seconds := time.Duration(1000 / 30)
	// // ticker := time.Tick(seconds * time.Millisecond)

}

// func (g *GameState) broadcastState(t <-chan time.Time) {

// 	for {
// 		select {
// 		case <-t:
// 			g.EventQueue.FireGameState(message.SendState())
// 		}

// 	}
// }

//GameState Whole game state
type GameState struct {
	requests    chan message.UserInput
	Players     map[uint32]*entity.Player
	Projectiles map[uint32]*entity.Projectile
	EventQueue  *events.EventQueue
}

//HandleInput request
func (g *GameState) HandleInput(m *message.NetworkInput) {

	player := g.Players[m.ID]
	player.UpdatePlayer(m)

	now := time.Now()
	before := player.LastShot

	diff := now.Sub(before) / time.Millisecond
	// println(diff)
	if player.IsShooting && diff >= 200 {
		player.LastShot = time.Now()
		newID := ProjectileID(10000, g.Projectiles)
		newProjectile := &entity.Projectile{
			Rotation: &types.Vector{
				X: player.Rotation.X,
				Y: player.Rotation.Y,
			},
			ID: (newID),
			Position: &types.Vector{
				X: float64(player.Position.X),
				Y: float64(player.Position.Y),
			},
			PlayerID: player.ID,
		}

		g.EventQueue.FireProjectileReady(newProjectile)
	}

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

//ProjectileID  Creates and Validates ID to be unique
func ProjectileID(size int, projectiles map[uint32]*entity.Projectile) uint32 {
	uniqueID := validate.GenerateID(size)
	if _, ok := projectiles[uniqueID]; ok {
		uniqueID = ProjectileID(size, projectiles)
	}
	return uniqueID
}

//HandleProjectileFired spawns a projecrtile into game state
func (g *GameState) HandleProjectileFired(m *message.SpawnProjectile) {
	//
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
