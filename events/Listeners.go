package events

import "sockets/message"

//InputListener triggered on player input from client
type InputListener interface {
	HandleRequest(*message.UserInput)
}

//broadcast inputRequest to all subscribers
func (l *InputRequest) process() {
	for _, listener := range l.subscribers {
		listener.HandleRequest(l.payload)
	}
}

//ConnectedListener triggered when new player connects
type ConnectedListener interface {
	HandleRequest(*message.Connect)
}

func (l *PlayerConnected) process() {
	for _, listener := range l.subscribers {
		listener.HandleRequest(l.payload)
	}
}

//DisconnectedListener triggered when player disconnects
type DisconnectedListener interface {
	HandleRequest(*message.Disconnect)
}

//broadcaster
func (l *PlayerDisconnected) process() {
	for _, listener := range l.subscribers {
		listener.HandleRequest(l.payload)
	}
}

//PlayerKilledListener triggered when a player is killed
type PlayerKilledListener interface {
	HandleRequest(*message.KillPlayer)
}

func (l *PlayerKilled) process() {
	for _, listener := range l.subscribers {
		listener.HandleRequest(l.payload)
	}
}

//ProjectileHitListener triggered when projectile collisions
type ProjectileHitListener interface {
	HandleRequest(*message.ProjectileHit)
}

func (l *ProjectileCollision) process() {
	for _, listener := range l.subscribers {
		listener.HandleRequest(l.payload)
	}
}
