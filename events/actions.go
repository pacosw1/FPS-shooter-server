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
	e.criticalQueue <- request
}

//FirePhysicsDone send a connect request to the event queue
func (e *EventQueue) FirePhysicsDone() {

	request := &PhysicsDone{
		subscribers: e.PhysicsDoneListeners,
	}
	e.criticalQueue <- request
}

//FireProjectileReady send a connect request to the event queue
func (e *EventQueue) FireProjectileReady(p *entity.Projectile) {

	request := &ProjectileReady{
		payload:     p,
		subscribers: e.ProjectileReadyListeners,
	}
	e.criticalQueue <- request
}

//FireGameState sends an event request to broadcast gameState to clients
func (e *EventQueue) FireGameState(s *entity.Broadcast) {

	request := &BroadcastState{
		payload:    s,
		subcribers: e.StateBroadcastListeners,
	}
	e.criticalQueue <- request
}

//FireDisconnect send a disconnect request to the event queue
func (e *EventQueue) FireDisconnect(m *message.Disconnect) {

	request := &PlayerDisconnect{
		payload:     m,
		subscribers: e.DisconnectListeners,
	}
	e.criticalQueue <- request
}

//FireInput send a disconnect request to the event queue
func (e *EventQueue) FireInput(m *message.NetworkInput) {
	request := &InputRequest{
		payload:     m,
		subscribers: e.InputListeners,
	}
	e.criticalQueue <- request
}
