package state

import (
	"math"
	"sockets/entity"
	"sockets/events"
	"sockets/message"
	pb "sockets/protobuf"
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
		EventQ:        e,
		pendingInputs: make(chan *message.NetworkInput, 1000000),
		DeltaTime:     0,
		Before:        time.Now(),
	}
}

//Start the broadcast timer
func (g *GameState) Start() {
	println("Game State Initialized")
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
	requests      chan message.UserInput
	Players       map[uint32]*entity.Player
	Projectiles   map[uint32]*entity.Projectile
	EventQ        *events.EventQueue
	Before        time.Time
	DeltaTime     float64
	pendingInputs chan *message.NetworkInput
}

//HandleInput request
func (g *GameState) HandleInput(m *message.NetworkInput) {

	g.pendingInputs <- m

}

//HandleTimeStep runs physics step
func (g *GameState) HandleTimeStep(frame int) {
	g.UpdatePlayers()
	g.UpdatePhysics(frame)

	before := g.Before
	now := time.Now().Sub(before)
	g.Before = time.Now()

	t := float64(now/time.Millisecond) / 1000.0
	g.DeltaTime = t

}

//UpdatePhysics s
func (g *GameState) UpdatePhysics(frame int) {
	g.updateShooting()
	g.UpdatePlayers()

	g.checkHits(g.Players)
	f := int(math.Floor(60 / 10))
	if frame%f == 0 {
		g.EventQ.FireGameState(g.CopyState())
	}

	g.updateProjectiles()

}

// func (g *GameState) HandleStartBroadcast() {
// 	g.EventQ.FireGameState(g.CopyState())
// }

//Broadcast state
func (g *GameState) Broadcast() {
	tick := time.Tick(100 * time.Millisecond)

	for {
		select {
		case <-tick:
			g.EventQ.FireGameState(g.CopyState())
		default:
			time.Sleep(5 * time.Millisecond)
		}

	}

}

//UpdatePlayers update players each server tick
func (g *GameState) UpdatePlayers() {
	applied := 0
	for len(g.pendingInputs) > 0 {
		input := <-g.pendingInputs

		player, found := g.Players[input.ID]

		if found {
			player.UpdatePlayer(input)
			g.updatePlayerState(player)
		}
		applied++
	}

}

func (g *GameState) updatePlayerState(player *entity.Player) {
	// player.Update(g.DeltaTime)
	before := player.LastShot
	now := time.Now()
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

		g.Projectiles[newID] = newProjectile

		// g.EventQ.FireProjectileReady(newProjectile)
	}
	player.UpdateMovement(300 * (1.0 / 60))

	g.Players[player.ID] = player

}

func (g *GameState) updateShooting() {
	for _, player := range g.Players {

		before := player.LastShot
		now := time.Now()
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

			g.Projectiles[newID] = newProjectile

			// g.EventQ.FireProjectileReady(newProjectile)
		}
	}
}

//CopyState copy state to proto
func (g *GameState) CopyState() *entity.Broadcast {

	ogPlayers := g.Players
	ogProject := g.Projectiles

	players := make(map[uint32]*pb.Player)
	projectiles := make(map[uint32]*pb.Projectile)

	for key, value := range ogPlayers {
		players[key] = value.ToProto()
	}

	for key, value := range ogProject {
		projectiles[key] = value.ToProto()
	}

	return &entity.Broadcast{
		Players:     players,
		Projectiles: projectiles,
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

func (g *GameState) updateProjectiles() {
	projectiles := g.Projectiles

	for ID := range projectiles {
		projectile := projectiles[ID]
		g.updateProjectile(projectile, ID)
	}

}

func (g *GameState) updateProjectile(projectile *entity.Projectile, ID uint32) {

	speed := 1200.0 * g.DeltaTime

	direction := projectile.Rotation.Normalize()
	velocity := direction.Dot(speed)

	projectile.Position.X += (math.Floor(velocity.X*100) / 100)
	projectile.Position.Y += (math.Floor(velocity.Y*100) / 100)

	x := projectile.Position.X
	y := projectile.Position.Y

	if (x > 2000 || x < -500) || (y > 2000 || y < -500) {
		delete(g.Projectiles, projectile.ID)
		return
	}

}

func (g *GameState) checkHit(playerID, projectileID uint32) {

	player := g.Players[playerID]
	if player.Dead == true {
		return
	}
	projectile := g.Projectiles[projectileID]

	pRadius := 10
	ppRadius := 30

	dx := float64(projectile.Position.X - float64(player.Position.X))
	dy := float64(projectile.Position.Y - float64(player.Position.Y))

	distance := math.Sqrt((dx * dx) + dy*dy)
	R := float64(pRadius + ppRadius)
	if distance <= R {
		player.Health -= 10
		// fireCollison && firePlayerDead events
		if player.Health <= 0 {
			player.Dead = true
			g.EventQ.FireDisconnect(message.DisconnectMessage(playerID))
		}
		delete(g.Projectiles, projectileID)
	}

}

//CheckHits Collision Detection
func (g *GameState) checkHits(players map[uint32]*entity.Player) {
	projectiles := g.Projectiles

	for player := range players {
		for projectile := range projectiles {
			if player != projectiles[projectile].PlayerID {
				g.checkHit(player, projectile)
			}
		}
		if g.Players[player].Dead {
			delete(g.Players, player)
		}
	}
}
