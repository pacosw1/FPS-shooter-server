package events

import (
	"sockets/entity"
	"sockets/message"
)

//FireConnect send a connect request to the event queue
func (e *EventQueue) FireConnect(m *message.Connect) {

	request := &PlayerConnect{
		payload:     m,
		subscribers: e.ConnectListeners,
	}
	e.primaryQ <- request
}

func (e *EventQueue) FireStartBroadcast() {

	request := &StartBroadcast{

		subscribers: e.StartBroadcastListeners,
	}
	e.primaryQ <- request
}

//FireTimeStep send a connect request to the event queue
func (e *EventQueue) FireTimeStep(frame int) {

	request := &TimeStep{
		payload:     frame,
		subscribers: e.TimeStepListeners,
	}
	e.criticalQ <- request
}

//FireProjectileReady send a connect request to the event queue
func (e *EventQueue) FireProjectileReady(p *entity.Projectile) {

	request := &ProjectileReady{
		payload:     p,
		subscribers: e.ProjectileReadyListeners,
	}
	e.primaryQ <- request
}

//FireGameState sends an event request to broadcast gameState to clients
func (e *EventQueue) FireGameState(s *entity.Broadcast) {

	request := &BroadcastState{
		payload:    s,
		subcribers: e.StateBroadcastListeners,
	}
	e.criticalQ <- request
}

//FireDisconnect send a disconnect request to the event queue
func (e *EventQueue) FireDisconnect(m *message.Disconnect) {

	request := &PlayerDisconnect{
		payload:     m,
		subscribers: e.DisconnectListeners,
	}
	e.primaryQ <- request
}

//FireInput send a disconnect request to the event queue
func (e *EventQueue) FireInput(m *message.NetworkInput) {
	request := &InputRequest{
		payload:     m,
		subscribers: e.InputListeners,
	}
	e.secondaryQ <- request
}
