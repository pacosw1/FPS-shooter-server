package entity

import (
	pb "sockets/protobuf"
	"sockets/types"
)

//Projectile stores bullet postion and angle (4 bytes)
type Projectile struct {
	Rotation *types.Vector
	Position *types.Vector
	ID       uint32
	PlayerID uint32
}

//ToProto converts
func (p *Projectile) ToProto() *pb.Projectile {
	return &pb.Projectile{
		Position: p.Position.ToProto(),
		// Rotation: p.Rotation.ToProto(),
	}
}
