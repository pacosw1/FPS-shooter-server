package object

//Player Stores state data for a player
type Player struct {
	Health int
	*Position
	Aim        int
	IsShooting bool
	SequenceID int16
	ID         int
}

//UserInput request to update player state
type UserInput struct {
	IsShooting bool
	// *Position
	// SequenceID int16
	// ID         int
}

//Projectile stores bullet postion and angle
type Projectile struct {
	Aim      *Position
	Position *Position
	ID       int
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
