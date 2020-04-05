package engine

import (
	"sockets/object"
	"time"
)

//GameLoop tick
func GameLoop(engine *Engine) {
	frames := time.Duration(1000 / engine.FPS)

	tick := time.Tick(frames * time.Millisecond)
	for engine.State == 1 {
		select {
		case <-tick:
			select {
			case handler := <-engine.requests:
				println(handler.IsShooting)
			default:
				println("waiting")
			}
		}

	}
}

//Start 's the game loop
func Start(e *Engine) {
	e.State = 1
	go GameLoop(e)
}

//HandleRequest action ran on event trigger
func (e *Engine) HandleRequest(payload *object.UserInput) {
	e.requests <- payload
}

//Stop stops the game loop from running
func Stop(e *Engine) {
	e.State = 0
}

//SetFPS refresh rate
func SetFPS(e *Engine, frames int) {
	e.FPS = frames
}

//Engine test
type Engine struct {
	requests chan *object.UserInput
	FPS      int
	// GameState  int

	State int
}

//New s
func New() *Engine {
	return &Engine{
		requests: make(chan *object.UserInput, 100),
		FPS:      1,
		State:    0,
	}
}
