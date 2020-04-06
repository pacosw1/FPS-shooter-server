package message

import "sockets/types"

//Connect request to create new Player on user joined
type Connect struct {
	Username string
	ClientID int
}

//ConnectMessage creates a new message of type Connect
func ConnectMessage(username string, ID int) *Connect {
	return &Connect{
		Username: username,
		ClientID: ID,
	}
}

//Disconnect request to disconnect player from server
type Disconnect struct {
	ClientID int
}

//DisconnectMessage creates a new message of type Disconnect
func DisconnectMessage(ID int) *Disconnect {
	return &Disconnect{
		ClientID: ID,
	}
}

//UserInput request to update player state
type UserInput struct {
	IsShooting bool
	*types.Position
	SequenceID int16
	Aim        *types.Position
	ID         int
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
