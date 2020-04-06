package message

import "sockets/types"

//Connect request to create new Player on user joined
type Connect struct {
	Username string
	ClientID int
}

//Disconnect request to disconnect player from server
type Disconnect struct {
	ClientID int
}

//UserInput request to update player state
type UserInput struct {
	IsShooting bool
	// *Position
	// SequenceID int16
	// ID         int
}

//SpawnProjectile message to spawn a projectile
type SpawnProjectile struct {
	ID       int16
	Position *types.Position
}

//KillPlayer event trigger on delete
type KillPlayer struct {
	ID int
}

//ProjectileHit message to destroy projectile
type ProjectileHit struct {
	ID int16
}
