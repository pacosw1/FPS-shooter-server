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

//FireDisconnect send a disconnect request to the event queue
func (e *EventQueue) FireDisconnect(m *message.Disconnect) {

	request := &PlayerDisconnect{
		payload:     m,
		subscribers: e.DisconnectListeners,
	}
	e.criticalQueue <- request
}

//FireInput send a disconnect request to the event queue
func (e *EventQueue) FireInput(m *message.UserInput) {

	request := &InputRequest{
		payload:     m,
		subscribers: e.InputListeners,
	}
	e.criticalQueue <- request
}
