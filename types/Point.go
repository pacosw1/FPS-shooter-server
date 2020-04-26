package types

import (
	pb "sockets/protobuf"
)

// Point Stores a 2D position
type Point struct {
	X float32
	Y float32
}

//ToProto pb point to proto
func (p *Point) ToProto() *pb.Point {
	return &pb.Point{
		X: float64(p.X),
		Y: float64(p.Y),
	}
}
