package entity

import (
	"math"
	"sockets/message"
	"sockets/types"
	"time"
)

// Player Stores state data for a player (12 bytes)
type Player struct {
	Health     uint8         //1
	Position   *types.Vector //4
	Rotation   *types.Vector //4
	IsShooting bool
	SequenceID uint16 //2
	ID         int    //1
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
	speed := 3
	p.SequenceID = uint16(r.SequenceID)
	p.IsShooting = r.IsShooting

	if r.Rotation != 0 {
		p.Rotate(r.Rotation)
	}

	// p.Position
	// p.Rotation.Normalize()
	if r.Direction != 0 {
		p.Position = p.Position.Add(p.Rotation.Normalize(), speed)

	}

}

//NewPlayer create a new player
func NewPlayer(clientID int) *Player {
	return &Player{
		Health: 100,
		Position: &types.Vector{
			X: 500,
			Y: 500,
		},
		Rotation: &types.Vector{
			X: 0,
			Y: -1,
		},
		IsShooting: false,
		SequenceID: 0,
		ID:         clientID,
		LastShot:   time.Now(),
		Dead:       false,
	}
}

//Rotate rotates character facing vector
func (p *Player) Rotate(d int) {

	degree := (1.0 / 15) * float64(d)
	X := float64(p.Rotation.X)
	Y := float64(p.Rotation.Y)
	dx := math.Cos(degree)*X - math.Sin(degree)*Y
	dy := math.Sin(degree)*X + math.Cos(degree)*Y

	p.Rotation.X = float32(dx)
	p.Rotation.Y = float32(dy)

}

//Projectile stores bullet postion and angle (4 bytes)
type Projectile struct {
	Rotation *types.Vector
	Position *types.Vector
	ID       int
	PlayerID int
}

//Zombie t
