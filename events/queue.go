package events

import (
	"sockets/object"
)

//EventQueue central structure to proccess all incoming client requests
type EventQueue struct {
	Running        bool
	criticalQueue  chan Request
	InputListeners []InputListener
}

//New EventQueue Instance
func New() *EventQueue {
	return &EventQueue{
		Running:        false,
		criticalQueue:  make(chan Request, 100),
		InputListeners: []InputListener{},
	}
}

//Start the event queue
func Start(e *EventQueue) {
	e.Running = true
	for e.Running {
		select {
		case request := <-e.criticalQueue:
			request.process()
		}

	}
}

//InputListener listen for inputs
type InputListener interface {
	HandleRequest(*object.UserInput)
}

//Request interface to send thru channel
type Request interface {
	process()
}

//InputRequest carries request payload thru channel
type InputRequest struct {
	payload     *object.UserInput
	subscribers []InputListener
}

//RegisterInput subscribe to listen User Input requests
func (q *EventQueue) RegisterInput(l InputListener) {
	q.InputListeners = append(q.InputListeners, l)
}

// func (l *InputRequest) broadcast() {
// 	println("broadcasting")
// }

//broadcast inputRequest to all subscribers
func (l *InputRequest) process() {
	for _, listener := range l.subscribers {
		listener.HandleRequest(l.payload)
	}
}

//SendInput sends input to event queue
func (q *EventQueue) SendInput(i *object.UserInput) {

	request := &InputRequest{
		payload:     i,
		subscribers: q.InputListeners,
	}
	q.criticalQueue <- request
}
