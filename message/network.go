package message

import (
	pb "sockets/protobuf"
	"sockets/types"
)

//NetworkMessage t
type NetworkMessage struct {
	Data string `json:"data"`
}

//NetworkInput (6 bytes)
type NetworkInput struct {
	IsShooting bool
	Direction  *types.Point
	SequenceID uint32
	Rotation   *types.Vector
	ID         uint32 `json:"ID"`
}

//ProtoPoint turns proto-point to Point
func ProtoPoint(p *pb.Point) *types.Point {
	return &types.Point{
		X: int32(p.X),
		Y: p.Y,
	}
}

//ProtoVector turns proto-vector to Vector
func ProtoVector(p *pb.Vector) *types.Vector {
	return &types.Vector{
		X: float64(p.X),
		Y: float64(p.Y),
	}
}

//UpdatePlayer <- updates player based on input
