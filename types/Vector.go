package types

import (
	"math"
	pb "sockets/protobuf"
)

// Point Stores a 2D position

//Vector 2D vector
type Vector struct {
	X float64
	Y float64
}

//Vector2D creates new vector
func Vector2D(x, y float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}

//ToProto t
func (v *Vector) ToProto() *pb.Vector {
	return &pb.Vector{
		X: float32(v.X),
		Y: float32(v.Y),
	}
}

//Length returns the length of vector v, (hypothenus)
func (v *Vector) Length() float64 {
	squared := (v.X * v.X) + (v.Y * v.Y)
	root := math.Sqrt(squared)
	return root

}

//Distance returnes the  distance between v and u
func (v *Vector) Distance(u *Vector) float64 {
	dx := (v.X - u.X)
	dy := (v.Y - u.Y)

	dist := math.Sqrt((dx * dx) + (dy * dy))
	return dist
}

//DistanceSq returnes the squared distance between v and u
func (v *Vector) DistanceSq(u *Vector) float64 {
	dx := (v.X - u.X)
	dy := (v.Y - u.Y)

	dist := ((dx * dx) + (dy * dy))
	return dist
}

//Dot returns the dot product between u * v
func (v *Vector) Dot(num float64) *Vector {
	return &Vector{
		X: v.X * num,
		Y: v.Y * num,
	}
}

//LengthSquared returns the squared value of hypothenus
func (v *Vector) LengthSquared() float64 {
	return (v.X * v.X) + (v.Y * v.Y)

}

//Inverse returns the inverse vector
func (v *Vector) Inverse() *Vector {
	return &Vector{
		X: -v.X,
		Y: -v.Y,
	}
}

//Normalize normalizes the vector by dividing by length
func (v *Vector) Normalize() *Vector {

	if v.X == 0 && v.Y == 0 {
		return &Vector{
			X: 0, Y: 0,
		}
	}
	mag := v.Length()
	x := (v.X / mag)
	y := (v.Y / mag)
	return &Vector{
		X: x,
		Y: y,
	}
}

//Add adds two vectors together
func (v *Vector) Add(u *Vector) *Vector {

	x := (v.X + u.X)
	y := (v.Y + u.Y)
	return &Vector{
		X: (x),
		Y: (y),
	}
}
