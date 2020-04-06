package simulation

import (
	"sockets/message"
	"time"
)

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

//Start 's the game loop
func (e *Engine) Start() {
	e.State = 1
	frames := time.Duration(1000 / e.FPS)
	tick := time.Tick(frames * time.Millisecond)
	go e.GameLoop(tick)
}

//Stop stops the game loop from running
func (e *Engine) Stop() {
	e.State = 0
}

//GameLoop tick
func (e *Engine) GameLoop(t <-chan time.Time) {
	for e.State == 1 {
		select {
		case <-t:
			select {
			case handler := <-e.requests:
				println(handler.IsShooting)
			default:
				println("waiting")
			}
		}

	}
}

//HandleRequest action ran on event trigger
func (e *Engine) HandleInput(payload *message.UserInput) {
	e.requests <- payload
}

//SetFPS refresh rate
func SetFPS(e *Engine, frames int) {
	e.FPS = frames
}
