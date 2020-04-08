package events

//EventQueue central structure to proccess all incoming client requests
type EventQueue struct {
	Running                 bool
	criticalQueue           chan Request
	InputListeners          []InputListener
	ConnectListeners        []ConnectListener
	DisconnectListeners     []DisconnectListener
	StateBroadcastListeners []StateBroadcastListener
}

//NewEventQ Instance
func NewEventQ() *EventQueue {
	return &EventQueue{
		Running:                 false,
		criticalQueue:           make(chan Request, 100),
		InputListeners:          []InputListener{},
		ConnectListeners:        []ConnectListener{},
		DisconnectListeners:     []DisconnectListener{},
		StateBroadcastListeners: []StateBroadcastListener{},
	}
}

//Start the event queue
func (e *EventQueue) Start() {
	e.Running = true
	go e.runLoop()

}

func (e *EventQueue) runLoop() {
	for e.Running {
		select {
		case request := <-e.criticalQueue:
			request.process()
			// default:
			// 	println("waiting")
		}

	}
}

//RegisterInput subscribe to listen User Input requests
func (e *EventQueue) RegisterInput(l InputListener) {
	e.InputListeners = append(e.InputListeners, l)
}

//RegisterConnect t
func (e *EventQueue) RegisterConnect(l ConnectListener) {
	e.ConnectListeners = append(e.ConnectListeners, l)
}

//RegisterBroadcast t
func (e *EventQueue) RegisterBroadcast(l StateBroadcastListener) {
	e.StateBroadcastListeners = append(e.StateBroadcastListeners, l)
}

//RegisterDisconnect t
func (e *EventQueue) RegisterDisconnect(l DisconnectListener) {
	e.DisconnectListeners = append(e.DisconnectListeners, l)
}
