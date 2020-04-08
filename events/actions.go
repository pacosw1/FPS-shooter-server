package events

import (
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

//FireGameState sends an event request to broadcast gameState to clients
func (e *EventQueue) FireGameState(m *message.StateMessage) {

	request := &BroadcastState{
		payload:    m,
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
