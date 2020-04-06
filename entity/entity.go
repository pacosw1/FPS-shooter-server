package entity

// Player Stores state data for a player
type Player struct {
	Health int
	*Position
	Aim        int
	IsShooting bool
	SequenceID int16
	ID         int
}

//Projectile stores bullet postion and angle
type Projectile struct {
	Aim      *Position
	Position *Position
	ID       int
}
