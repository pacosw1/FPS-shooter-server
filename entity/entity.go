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
	}
}

//UpdatePlayer <- updates player based on input
func (p *Player) UpdatePlayer(r *message.NetworkInput) {
	p.Position.X += r.Direction.X * 5
	p.Position.Y += r.Direction.Y * 5
	p.SequenceID = r.SequenceID
	p.IsShooting = r.IsShooting
	p.Aim = r.Aim
}

//Projectile stores bullet postion and angle
type Projectile struct {
	Aim      *types.Position
	Position *types.Position
	ID       int
}
