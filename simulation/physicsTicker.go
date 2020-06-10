package simulation

import (
	"sockets/events"
	"time"
)

//PhysicsTicker t
type PhysicsTicker struct {
	fps     int
	eventQ  *events.EventQueue
	FrameID uint32
}

//NewPhysicsTicker create ticker
func NewPhysicsTicker(q *events.EventQueue) *PhysicsTicker {
	return &PhysicsTicker{
		fps:     60,
		eventQ:  q,
		FrameID: 0,
	}
}

//Run start the ticker
func (ticker *PhysicsTicker) Run() {
	i := 0
	x := (1000 / 60)
	br := time.Duration(x)
	fps := time.Millisecond * br
	for range time.Tick(fps) {
		ticker.eventQ.FireTimeStep(i)
		i++
	}
}
