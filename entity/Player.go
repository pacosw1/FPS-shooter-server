package entity

import (
	"math"
	"sockets/message"
	pb "sockets/protobuf"
	"sockets/types"
	"time"
)

// Player Stores state data for a player (12 bytes)
type Player struct {
	Health     uint32       //1
	Position   *types.Point //8
	Direction  *types.Point
	Rotation   *types.Vector //8
	IsShooting bool
	SequenceID uint32 //2
	ID         uint32 //1
	LastShot   time.Time
	Dead       bool
}

//ToProto turns buffer
func (p *Player) ToProto() *pb.Player {
	return &pb.Player{
		RequestID: int32(p.SequenceID),
		Position:  p.Position.ToProto(),
		Rotation:  p.Rotation.ToProto(),
		Hp:        int32(p.Health),
	}

}

//UpdatePlayer t
func (p *Player) UpdatePlayer(r *message.NetworkInput) {
	p.SequenceID = (r.SequenceID) //check overload of uint 16
	p.IsShooting = r.IsShooting
	p.Direction = r.Direction
	p.Rotation = r.Rotation

}

//NewPlayer create a new player
func NewPlayer(clientID uint32) *Player {
	return &Player{
		Health: 100,
		Position: &types.Point{
			X: 500,
			Y: 500,
		},
		Direction: &types.Point{
			X: 0,
			Y: 0,
		},
		Rotation: &types.Vector{
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

//Update updates player movement each frame;
func (p *Player) Update(dt float64) {
	// p.SequenceID++
	p.UpdateMovement(dt)
}

//UpdateMovement t
func (p *Player) UpdateMovement(dt float64) {

	dir := p.Direction
	move := &types.Vector{
		X: float64(dir.X),
		Y: float64(dir.Y),
	}

	speed := float64(300.0 * dt)
	// fmt.Println(dt)
	move = move.Dot(speed)
	p.Position.X = int32(math.Round((move.X + float64(p.Position.X))))
	p.Position.Y = int32(math.Round((move.Y + float64(p.Position.Y))))

	// println(p.Position.X)

	// //normalize rotation
	// rotation := p.Rotation.Normalize()

	// //direction 0 = idle, 1 = forward, -1 = backward   * velocity
	// direction := rotation.Dot(d * v)

	// //add vectors
	// p.Position = p.Position.Add(direction)

}

//UpdateRotation rotates character facing vector
func (p *Player) UpdateRotation() {

	//if player isn't rotating exit function

	// fmt.Println(x)
	// if d == 0 {
	// 	return
	// }

	// //else calculate angle and updar
	// m := float64(d)
	// degree := float64(0.1 * m)
	// X := float64(p.Rotation.X)
	// Y := float64(p.Rotation.Y)
	// dx := math.Cos(degree)*X - math.Sin(degree)*Y
	// dy := math.Sin(degree)*X + math.Cos(degree)*Y

	// p.Rotation.X = float32(dx)
	// p.Rotation.Y = float32(dy)

}
