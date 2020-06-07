package simulation

import (
	"math"
	"sockets/entity"
	"sockets/events"
	"sockets/message"
	"sockets/state"
	"sockets/types"
	"sockets/utils"
	"sockets/validate"
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

func (e *Engine) copyPlayers() map[uint32]*entity.Player {
	original := e.GameState.Players
	copy := make(map[uint32]*entity.Player)

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
	br := time.Duration(1000.0 / 10)
	tick := time.Tick(br * time.Millisecond)
	for e.State == 1 {
		select {
		case <-t:
			e.updatePhysics()
		case <-tick:
			e.Broadcast(e.GameState)
		}

	}
}

func (e *Engine) updatePlayers(players map[uint32]*entity.Player) {

	for ID := range players {
		e.updatePlayer(ID)
	}
}

func (e *Engine) updatePhysics() {
	players := e.copyPlayers()
	e.updatePlayers(players)
	e.loadProjectiles()
	e.checkHits(players)
	e.updateProjectiles()
}

//ProjectileID  Creates and Validates ID to be unique
func ProjectileID(size int, projectiles map[uint32]*entity.Projectile) uint32 {
	uniqueID := validate.GenerateID(size)
	if _, ok := projectiles[uniqueID]; ok {
		uniqueID = ProjectileID(size, projectiles)
	}
	return uniqueID
}

func (e *Engine) updatePlayer(ID uint32) {

	p := e.GameState.Players[ID]

	speed := 3

	p.UpdateMovement(speed)
	if (p.Direction.X != 0 && p.Direction.Y != 0) || !p.IsShooting {
		p.SequenceID++
	}
	before := p.LastShot
	now := time.Now()

	diff := now.Sub(before) / time.Millisecond
	// println(diff)
	if p.IsShooting && diff >= 200 {
		p.LastShot = time.Now()
		newID := ProjectileID(10000, e.GameState.Projectiles)
		newProjectile := &entity.Projectile{
			Rotation: &types.Vector{
				X: p.Rotation.X,
				Y: p.Rotation.Y,
			},
			ID: (newID),
			Position: &types.Vector{
				X: float64(p.Position.X),
				Y: float64(p.Position.Y),
			},
			PlayerID: p.ID,
		}
		//adds projectiles to queue
		e.pendingProjectiles <- newProjectile
	}
}

func (e *Engine) updateProjectiles() {
	projectiles := e.GameState.Projectiles

	for ID := range projectiles {
		projectile := projectiles[ID]
		e.updateProjectile(projectile, ID)
	}

}

func (e *Engine) updateProjectile(projectile *entity.Projectile, ID uint32) {

	speed := 5

	direction := projectile.Rotation.Normalize()
	velocity := direction.Dot(speed)

	projectile.Position.X += (math.Floor(velocity.X*100) / 100)
	projectile.Position.Y += (math.Floor(velocity.Y*100) / 100)

	x := projectile.Position.X
	y := projectile.Position.Y

	if (x > 2000 || x < -500) || (y > 2000 || y < -500) {
		delete(e.GameState.Projectiles, projectile.ID)
		return
	}

}

func (e *Engine) checkHit(playerID, projectileID uint32) {

	player := e.GameState.Players[playerID]
	if player.Dead == true {
		return
	}
	projectile := e.GameState.Projectiles[projectileID]

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
			e.EventQ.FireDisconnect(message.DisconnectMessage(playerID))
		}
		delete(e.GameState.Projectiles, projectileID)
	}

}

//CheckHits Collision Detection
func (e *Engine) checkHits(players map[uint32]*entity.Player) {
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
