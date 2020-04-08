package events

//broadcaster
func (l *PlayerDisconnect) process() {
	for _, listener := range l.subscribers {
		listener.HandleDisconnect(l.payload)
	}
}

//broadcast inputRequest to all subscribers
func (l *InputRequest) process() {
	for _, listener := range l.subscribers {
		listener.HandleInput(l.payload)
	}
}

func (l *PlayerConnect) process() {
	for _, listener := range l.subscribers {
		listener.HandleConnect(l.payload)
	}
}

func (l *BroadcastState) process() {
	for _, listener := range l.subcribers {
		listener.HandleStateBroadcast(l.payload)
	}
}

func (l *PlayerKilled) process() {
	for _, listener := range l.subscribers {
		listener.handlePlayerKill(l.payload)
	}
}

func (l *ProjectileCollision) process() {
	for _, listener := range l.subscribers {
		listener.handleProjectileHit(l.payload)
	}
}
