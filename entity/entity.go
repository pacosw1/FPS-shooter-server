package entity

import (
	"fmt"
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
	speed := 5
	p.SequenceID = uint16(r.SequenceID)
	p.IsShooting = r.IsShooting

	//update player position and facing vector (rotation)
	p.UpdateRotation(r.Rotate)
	p.UpdateMovement(r.Direction, speed)

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

//UpdateMovement t
func (p *Player) UpdateMovement(d, v int) {

	//normalize rotation
	rotation := p.Rotation.Normalize()

	//direction 0 = idle, 1 = forward, -1 = backward   * velocity
	direction := rotation.Dot(d * v)

	//add vectors
	p.Position = p.Position.Add(direction)

}

//UpdateRotation rotates character facing vector
func (p *Player) UpdateRotation(d int) {

	//if player isn't rotating exit function
	if d == 0 {
		return
	}

	//else calculate angle and updar
	m := float64(d)
	degree := float64(0.05 * m)
	X := float64(p.Rotation.X)
	Y := float64(p.Rotation.Y)
	dx := math.Cos(degree)*X - math.Sin(degree)*Y
	dy := math.Sin(degree)*X + math.Cos(degree)*Y

	p.Rotation.X = float32(dx)
	p.Rotation.Y = float32(dy)
	fmt.Println(p.Rotation)

}

//Projectile stores bullet postion and angle (4 bytes)
type Projectile struct {
	Rotation *types.Vector
	Position *types.Vector
	ID       int
	PlayerID int
}

//Zombie t
