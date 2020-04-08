package events

import "sockets/message"

//InputListener triggered on player input from client
type InputListener interface {
	HandleInput(*message.NetworkInput)
}

//ConnectListener triggered when new player connects
type ConnectListener interface {
	HandleConnect(*message.Connect)
}

//StateBroadcastListener listens on state broadcast to clients
type StateBroadcastListener interface {
	HandleStateBroadcast(*message.StateMessage)
}

//PlayerKillListener triggered when a player is killed
type PlayerKillListener interface {
	handlePlayerKill(*message.KillPlayer)
}

//ProjectileHitListener triggered when projectile collisions
type ProjectileHitListener interface {
	handleProjectileHit(*message.ProjectileHit)
}

//DisconnectListener triggered when player disconnects
type DisconnectListener interface {
	HandleDisconnect(*message.Disconnect)
}
