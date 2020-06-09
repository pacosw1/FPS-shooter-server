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

func (l *TimeStep) process() {
	for _, listener := range l.subscribers {
		listener.HandleTimeStep(l.payload)
	}
}

func (l *StartBroadcast) process() {
	for _, listener := range l.subscribers {
		listener.HandleStartBroadcast()
	}
}

//projectile ready
func (l *ProjectileReady) process() {
	for _, listener := range l.subscribers {
		listener.HandleProjectileReady(l.payload)
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
