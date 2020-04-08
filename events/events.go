package events

import "sockets/message"

//Request interface to send different requests through the same channel
type Request interface {
	process()
}

//InputRequest carries request payload thru channel
type InputRequest struct {
	payload     *message.NetworkInput
	subscribers []InputListener
}

//BroadcastState event
type BroadcastState struct {
	payload    *message.StateMessage
	subcribers []StateBroadcastListener
}

//PlayerDisconnect event
type PlayerDisconnect struct {
	payload     *message.Disconnect
	subscribers []DisconnectListener
}

//ProjectileCollision event
type ProjectileCollision struct {
	payload     *message.ProjectileHit
	subscribers []ProjectileHitListener
}

//PlayerKilled event
type PlayerKilled struct {
	payload     *message.KillPlayer
	subscribers []PlayerKillListener
}

//PlayerConnect event fired when a new player connects
type PlayerConnect struct {
	payload     *message.Connect
	subscribers []ConnectListener
}
