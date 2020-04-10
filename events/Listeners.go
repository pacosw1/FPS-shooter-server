package events

import (
	"sockets/entity"
	"sockets/message"
)

//InputListener triggered on player input from client
type InputListener interface {
	HandleInput(*message.NetworkInput)
}

//ConnectListener triggered when new player connects
type ConnectListener interface {
	HandleConnect(*message.Connect)
}

//PhysicsDoneListener triggered when physics simulation step done
type PhysicsDoneListener interface {
	HandlePhysicsDone()
}

//StateBroadcastListener listens on state broadcast to clients
type StateBroadcastListener interface {
	HandleStateBroadcast(*entity.Broadcast)
}

//PlayerKillListener triggered when a player is killed
type PlayerKillListener interface {
	handlePlayerKill(*message.KillPlayer)
}

//ProjectileReadyListener triggered when projectile fired
type ProjectileReadyListener interface {
	HandleProjectileReady(*entity.Projectile)
}

//ProjectileHitListener triggered when projectile collisions
type ProjectileHitListener interface {
	handleProjectileHit(*message.ProjectileHit)
}

//DisconnectListener triggered when player disconnects
type DisconnectListener interface {
	HandleDisconnect(*message.Disconnect)
}
