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
		fps:     30,
		eventQ:  q,
		FrameID: 0,
	}
}

//Run start the ticker
func (ticker *PhysicsTicker) Run() {
	i := 0
	br := time.Duration(1000 / 30)
	fps := time.Millisecond * br
	for range time.Tick(fps) {
		ticker.eventQ.FireTimeStep(i)
		i++
	}
}
