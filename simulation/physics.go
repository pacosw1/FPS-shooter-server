package simulation

import (
	"math"
	"sockets/entity"
	"sockets/events"
	"sockets/message"
	"sockets/state"
	"sockets/utils"
	"time"
)

//Engine test
type Engine struct {
	pendingProjectiles chan *entity.Projectile
	EventQ             *events.EventQueue
	GameState          *state.GameState
	FPS                int
	State              int
}

//New creates a new Simulation Engine
func New(g *state.GameState, e *events.EventQueue) *Engine {
	return &Engine{
		pendingProjectiles: make(chan *entity.Projectile, 1000000),
		EventQ:             e,
		GameState:          g,
		FPS:                60,
		State:              0,
	}
}

func (e *Engine) loadProjectiles() {
	for len(e.pendingProjectiles) > 0 {
		projectile := <-e.pendingProjectiles
		e.GameState.Projectiles[projectile.ID] = projectile
	}
}

//Start 's the game loop
func (e *Engine) Start() {
	println("Physics Simulation Online")
	e.State = 1
	frames := time.Duration(1000 / e.FPS)
	tick := time.Tick(frames * time.Millisecond)
	go e.GameLoop(tick)
}

//Stop stops the game loop from running
func (e *Engine) Stop() {
	e.State = 0
}

func (e *Engine) copyPlayers() map[int]*entity.Player {
	original := e.GameState.Players
	copy := make(map[int]*entity.Player)

	for key, value := range original {
		copy[key] = value
	}
	return copy
}

//HandleProjectileReady adds new projectile to the queue
func (e *Engine) HandleProjectileReady(p *entity.Projectile) {
	e.pendingProjectiles <- p
}

//Broadcast copies state to send to client prevents memory access
func (e *Engine) Broadcast(s *state.GameState) {
	copy := utils.CopyState(s)
	e.EventQ.FireGameState(copy)
}

//GameLoop tick
func (e *Engine) GameLoop(t <-chan time.Time) {
	// br := time.Duration(1000 / 10)
	// tick := time.Tick(br * time.Millisecond)
	for e.State == 1 {
		select {
		case <-t:
			players := e.copyPlayers()
			e.loadProjectiles()

			e.checkHits(players)
			e.updateProjectiles()
			e.Broadcast(e.GameState)
		}

	}
}

func (e *Engine) updateProjectiles() {
	projectiles := e.GameState.Projectiles

	for ID := range projectiles {
		projectile := projectiles[ID]
		e.updateProjectile(projectile, ID)
	}

}

func (e *Engine) updateProjectile(projectile *entity.Projectile, ID int) {

	// speed := 10

	projectile.Rotation.Normalize()
	projectile.Position = projectile.Position.Add(projectile.Rotation.Normalize())

	x := projectile.Position.X
	y := projectile.Position.Y

	if (x > 2000 || x < -500) || (y > 2000 || y < -500) {
		delete(e.GameState.Projectiles, projectile.ID)
		return
	}

}

func (e *Engine) checkHit(playerID, projectileID int) {

	player := e.GameState.Players[playerID]
	if player.Dead == true {
		return
	}
	projectile := e.GameState.Projectiles[projectileID]

	pRadius := 10
	ppRadius := 30

	dx := float64(projectile.Position.X - player.Position.X)
	dy := float64(projectile.Position.Y - player.Position.Y)

	distance := math.Sqrt((dx * dx) + dy*dy)
	R := float64(pRadius + ppRadius)
	if distance <= R {
		player.Health -= 10
		// fireCollison && firePlayerDead events
		if player.Health <= 0 {
			player.Dead = true
			e.EventQ.FireDisconnect(message.DisconnectMessage(playerID))
		}
		delete(e.GameState.Projectiles, projectileID)
	}

}

//CheckHits Collision Detection
func (e *Engine) checkHits(players map[int]*entity.Player) {
	projectiles := e.GameState.Projectiles

	for player := range players {
		for projectile := range projectiles {
			if player != projectiles[projectile].PlayerID {
				e.checkHit(player, projectile)
			}
		}
		if e.GameState.Players[player].Dead {
			delete(e.GameState.Players, player)
		}
	}
}
