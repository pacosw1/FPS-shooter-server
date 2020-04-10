package simulation

import (
	"math"
	"sockets/entity"
	"sockets/events"
	"sockets/state"
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
		pendingProjectiles: make(chan *entity.Projectile, 100),
		EventQ:             e,
		GameState:          g,
		FPS:                60,
		State:              0,
	}
}

func (e *Engine) loadProjectiles() {
	for len(e.pendingProjectiles) > 0 {
		projectile := <-e.pendingProjectiles
		println("loaded projectile")
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
	println("adding projectile to queue")
	e.pendingProjectiles <- p
}

//GameLoop tick
func (e *Engine) GameLoop(t <-chan time.Time) {
	for e.State == 1 {
		select {
		case <-t:
			players := e.copyPlayers()
			e.loadProjectiles()
			e.checkHits(players)
			e.updateProjectiles()

		}

	}
}

func (e *Engine) updateProjectiles() {
	projectiles := e.GameState.Projectiles

	for ID := range projectiles {
		projectile := projectiles[ID]
		e.updateProjectile(projectile)
	}

}

func (e *Engine) updateProjectile(projectile *entity.Projectile) {

	speed := 5

	projectile.Position.X += speed * projectile.Direction.X
	projectile.Position.Y += speed * projectile.Direction.Y

	x := projectile.Position.X
	y := projectile.Position.Y

	if (x > 2000 || x < 0) || (y > 2000 || y < 0) {
		delete(e.GameState.Projectiles, projectile.ID)
		println(projectile.Direction.X)
		println(projectile.Direction.Y)
		return
	}

	println("x:", x, " aimX: ", projectile.Direction.X, " aimY: ", projectile.Direction.Y)

}

func (e *Engine) checkHit(playerID, projectileID int) {
	player := e.GameState.Players[playerID]
	projectile := e.GameState.Projectiles[projectileID]

	pRadius := 10
	ppRadius := 30

	dx := float64(projectile.Position.X - player.Position.X)
	dy := float64(projectile.Position.Y - player.Position.Y)

	distance := math.Sqrt((dx * dx) + dy*dy)
	R := float64(pRadius + ppRadius)
	if distance <= R {
		player.Health -= 10
		//fireCollison && firePlayerDead events
		// if player.Health <= 0 {
		// 	e.GameState.RemovePlayer(message.DisconnectMessage(player))
		// }
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
	}
}
