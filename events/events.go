package events

import "sockets/message"

//Request interface to send different requests through the same channel
type Request interface {
	process()
}

//InputRequest carries request payload thru channel
type InputRequest struct {
	payload     *message.UserInput
	subscribers []InputListener
}

//PlayerDisconnected event
type PlayerDisconnected struct {
	payload     *message.Disconnect
	subscribers []DisconnectedListener
}

//ProjectileCollision event
type ProjectileCollision struct {
	payload     *message.ProjectileHit
	subscribers []ProjectileHitListener
}

//PlayerKilled event
type PlayerKilled struct {
	payload     *message.KillPlayer
	subscribers []PlayerKilledListener
}

//PlayerConnected event fired when a new player connects
type PlayerConnected struct {
	payload     *message.Connect
	subscribers []ConnectedListener
}
