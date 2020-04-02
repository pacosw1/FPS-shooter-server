package validation

import (
	"math/rand"
	object "sockets/object"
	"time"
)

// ValidatePlayerID Creates and Validates ID to be unique
func ValidatePlayerID(size int, players map[int]*object.Player) int {
	uniqueID := generateID(size)
	if _, ok := players[uniqueID]; ok {
		uniqueID = ValidatePlayerID(size, players)
	}
	return uniqueID
}

//ValidateProjectileID  Creates and Validates ID to be unique
func ValidateProjectileID(size int, projectiles map[int]*object.Projectile) int {
	uniqueID := generateID(size)
	if _, ok := projectiles[uniqueID]; ok {
		uniqueID = ValidateProjectileID(size, projectiles)
	}
	return uniqueID
}

//generaates a unique ID beteen 0 and 100
func generateID(size int) int {
	timestamp := time.Now().UnixNano()
	rand.Seed(timestamp)
	id := rand.Intn(size)
	return id
}
