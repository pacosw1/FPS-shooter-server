package events

//EventQueue central structure to proccess all incoming client requests
type EventQueue struct {
	Running        bool
	criticalQueue  chan Request
	InputListeners []InputListener
}

//NewEventQ Instance
func NewEventQ() *EventQueue {
	return &EventQueue{
		Running:        false,
		criticalQueue:  make(chan Request, 100),
		InputListeners: []InputListener{},
	}
}

//Start the event queue
func (e *EventQueue) Start() {
	e.Running = true
	for e.Running {
		select {
		case request := <-e.criticalQueue:
			request.process()
		}

	}
}

//RegisterInput subscribe to listen User Input requests
func (e *EventQueue) RegisterInput(l InputListener) {
	e.InputListeners = append(e.InputListeners, l)
}
