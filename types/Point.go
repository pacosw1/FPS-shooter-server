package types

import (
	pb "sockets/protobuf"
)

// Point Stores a 2D position
type Point struct {
	X int32
	Y int32
}

//ToProto pb point to proto
func (p *Point) ToProto() *pb.Point {
	return &pb.Point{
		X: (p.X),
		Y: (p.Y),
	}
}
