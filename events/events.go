package events

import (
	"sockets/entity"
	"sockets/message"
)

//Request interface to send different requests through the same channel
type Request interface {
	process()
}

//InputRequest carries request payload thru channel
type InputRequest struct {
	payload     *message.NetworkInput
	subscribers []InputListener
}

//TimeStep t
type TimeStep struct {
	payload     int
	subscribers []TimeStepListener
}

type StartBroadcast struct {
	subscribers []StartBroadcastListener
}

// //PhysicsDone carries request payload thru channel
// type PhysicsDone struct {
// 	subscribers []PhysicsDoneListener
// }

//BroadcastState event
type BroadcastState struct {
	payload    *entity.Broadcast
	subcribers []StateBroadcastListener
}

//PlayerDisconnect event
type PlayerDisconnect struct {
	payload     *message.Disconnect
	subscribers []DisconnectListener
}

//ProjectileReady event
type ProjectileReady struct {
	payload     *entity.Projectile
	subscribers []ProjectileReadyListener
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
