package object

//Player Stores state data for a player
type Player struct {
	Health int
	*Position
	SequenceID int16
	ID         int
}

//Projectile stores bullet postion and angle
type Projectile struct {
	Angle    *Angle
	Position *Position
	ID       float32
}

//Position Stores a 2D position
type Position struct {
	X int
	Y int
}

//Angle Position Angle for a 2D position
type Angle struct {
	X float64
	Y float64
}
