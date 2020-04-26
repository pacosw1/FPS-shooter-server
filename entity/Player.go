package entity

import (
	"sockets/message"
	pb "sockets/protobuf"
	"sockets/types"
	"time"
)

// Player Stores state data for a player (12 bytes)
type Player struct {
	Health     uint32        //1
	Position   *types.Point  //8
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
	speed := 3
	p.SequenceID = (r.SequenceID) //check overload of uint 16
	p.IsShooting = r.IsShooting
	//
	//update player position and facing vector (rotation)
	p.UpdateRotation(r.Rotation.X, r.Rotation.Y)
	p.UpdateMovement(r.Direction, speed)

}

//NewPlayer create a new player
func NewPlayer(clientID uint32) *Player {
	return &Player{
		Health: 100,
		Position: &types.Point{
			X: 500,
			Y: 500,
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

//UpdateMovement t
func (p *Player) UpdateMovement(dir *types.Point, v int) {

	move := &types.Vector{
		X: float64(dir.X),
		Y: float64(dir.Y),
	}
	move = move.Dot(v)
	p.Position.X += float32(move.X)
	p.Position.Y += float32(move.Y)

	// //normalize rotation
	// rotation := p.Rotation.Normalize()

	// //direction 0 = idle, 1 = forward, -1 = backward   * velocity
	// direction := rotation.Dot(d * v)

	// //add vectors
	// p.Position = p.Position.Add(direction)

}

//UpdateRotation rotates character facing vector
func (p *Player) UpdateRotation(x, y float64) {

	//if player isn't rotating exit function

	// fmt.Println(x)

	p.Rotation.X = x
	p.Rotation.Y = y
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
