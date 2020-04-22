package types

import "math"

// Direction Stores a 2D position
type Direction struct {
	X int
	Y int
}

//Vector 2D vector
type Vector struct {
	X float32
	Y float32
}

//Vector2D creates new vector
func Vector2D(x, y float32) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}

//Length returns the length of vector v, (hypothenus)
func (v *Vector) Length() float32 {
	squared := float64((v.X * v.X) + (v.Y * v.Y))
	root := math.Sqrt(squared)
	return float32(root)

}

//Distance returnes the  distance between v and u
func (v *Vector) Distance(u *Vector) float32 {
	dx := float64(v.X - u.X)
	dy := float64(v.Y - u.Y)

	dist := math.Sqrt((dx * dx) + (dy * dy))
	return float32(dist)
}

//DistanceSq returnes the squared distance between v and u
func (v *Vector) DistanceSq(u *Vector) float32 {
	dx := float32(v.X - u.X)
	dy := float32(v.Y - u.Y)

	dist := ((dx * dx) + (dy * dy))
	return dist
}

//Dot returns the dot product between u * v
func (v *Vector) Dot(m int) *Vector {
	num := float32(m)
	return &Vector{
		X: v.X * num,
		Y: v.Y * num,
	}

}

//LengthSquared returns the squared value of hypothenus
func (v *Vector) LengthSquared() float32 {
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
	x := float64(v.X/mag) * 100.0
	y := float64(v.Y/mag) * 100.0
	return &Vector{
		X: float32(x / 100.0),
		Y: float32(y / 100.0),
	}
}

//Add adds two vectors together
func (v *Vector) Add(u *Vector) *Vector {

	x := float64(v.X + u.X)
	y := float64(v.Y + u.Y)
	return &Vector{
		X: float32(math.Floor(x*100) / 100),
		Y: float32(math.Floor(y*100) / 100),
	}
}
