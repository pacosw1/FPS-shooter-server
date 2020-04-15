package entity

import (
	"sockets/message"
	"sockets/types"
	"time"
)

// Player Stores state data for a player
type Player struct {
	Health     int
	Position   *types.Position
	Aim        *types.Position
	IsShooting bool
	SequenceID int16
	ID         int
	LastShot   time.Time
	Dead       bool
}

//Broadcast s
type Broadcast struct {
	Players     map[int]*Player
	Projectiles map[int]*Projectile
}

//UpdatePlayer t
func (p *Player) UpdatePlayer(r *message.NetworkInput) {
	speed := 7
	p.Position.X += r.Direction.X * speed
	p.Position.Y += r.Direction.Y * speed
	p.SequenceID = r.SequenceID
	p.IsShooting = r.IsShooting
	p.Aim = r.Aim
}

//NewPlayer create a new player
func NewPlayer(clientID int) *Player {
	return &Player{
		Health: 100,
		Position: &types.Position{
			X: 0,
			Y: 0,
		},
		Aim: &types.Position{
			X: 0,
			Y: 0,
		},
		IsShooting: false,
		SequenceID: 0,
		ID:         clientID,
		LastShot:   time.Now(),
		Dead:       false,
	}
}

//Projectile stores bullet postion and angle
type Projectile struct {
	Direction *types.Position
	Position  *types.Position
	ID        int
	PlayerID  int
}

//Zombie t
type Zombie struct {
	Position *types.Position
	Angle    *types.Angle
	Target   *types.Position
	ID       int
}
