package simulation

import (
	"sockets/message"
	"time"
)

//GameLoop tick
func (e *Engine) GameLoop() {
	frames := time.Duration(1000 / e.FPS)

	tick := time.Tick(frames * time.Millisecond)
	for e.State == 1 {
		select {
		case <-tick:
			select {
			case handler := <-e.requests:
				println(handler.IsShooting)
			default:
				println("waiting")
			}
		}

	}
}

//Start 's the game loop
func (e *Engine) Start() {
	e.State = 1
	go e.GameLoop()
}

//HandleRequest action ran on event trigger
func (e *Engine) HandleRequest(payload *message.UserInput) {
	e.requests <- payload
}

//Stop stops the game loop from running
func (e *Engine) Stop() {
	e.State = 0
}

//SetFPS refresh rate
func SetFPS(e *Engine, frames int) {
	e.FPS = frames
}

//Engine test
type Engine struct {
	requests chan *message.UserInput
	FPS      int
	// GameState  int

	State int
}

//NewEngine creates a new Simulation Engine
func NewEngine() *Engine {
	return &Engine{
		requests: make(chan *message.UserInput, 100),
		FPS:      1,
		State:    0,
	}
}
