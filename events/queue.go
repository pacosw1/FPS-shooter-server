package events

import (
	"time"
)

//EventQueue central structure to proccess all incoming client requests
type EventQueue struct {
	Running                  bool
	criticalQ                chan Request
	primaryQ                 chan Request
	secondaryQ               chan Request
	InputListeners           []InputListener
	ConnectListeners         []ConnectListener
	DisconnectListeners      []DisconnectListener
	StateBroadcastListeners  []StateBroadcastListener
	StartBroadcastListeners  []StartBroadcastListener
	ProjectileReadyListeners []ProjectileReadyListener
	TimeStepListeners        []TimeStepListener
}

//NewEventQ Instance
func NewEventQ() *EventQueue {
	return &EventQueue{
		Running:    false,
		criticalQ:  make(chan Request, 1000000),
		primaryQ:   make(chan Request, 1000000),
		secondaryQ: make(chan Request, 1000000),

		InputListeners:           []InputListener{},
		ConnectListeners:         []ConnectListener{},
		DisconnectListeners:      []DisconnectListener{},
		StateBroadcastListeners:  []StateBroadcastListener{},
		StartBroadcastListeners:  []StartBroadcastListener{},
		ProjectileReadyListeners: []ProjectileReadyListener{},
		TimeStepListeners:        []TimeStepListener{},
	}
}

//Start the event queue
func (e *EventQueue) Start() {
	println("Event Queue Online")
	e.Running = true
	go e.runLoop()

}

func (e *EventQueue) runLoop() {

	for e.Running {
		select {
		case request := <-e.criticalQ:
			request.process()
		case request := <-e.primaryQ:
			request.process()
		case request := <-e.secondaryQ:
			request.process()
		default:
			time.Sleep(10 * time.Millisecond)

		}
	}
}

//RegisterInput subscribe to listen User Input requests
func (e *EventQueue) RegisterInput(l InputListener) {
	e.InputListeners = append(e.InputListeners, l)
}

func (e *EventQueue) RegisterStartBroadcast(l StartBroadcastListener) {
	e.StartBroadcastListeners = append(e.StartBroadcastListeners, l)
}

//RegisterTimeStep su
func (e *EventQueue) RegisterTimeStep(l TimeStepListener) {
	e.TimeStepListeners = append(e.TimeStepListeners, l)
}

//RegisterProjectileReady t
func (e *EventQueue) RegisterProjectileReady(l ProjectileReadyListener) {
	e.ProjectileReadyListeners = append(e.ProjectileReadyListeners, l)
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
