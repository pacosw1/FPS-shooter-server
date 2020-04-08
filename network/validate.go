package network

import (
	"math/rand"
	"sockets/entity"
	"time"
)

// PlayerID Creates and Validates ID to be unique
func PlayerID(size int, n *Network) int {
	uniqueID := generateID(size)
	if _, ok := n.Clients[uniqueID]; ok {
		uniqueID = PlayerID(size, n)
	}
	return uniqueID
}

//ValidateProjectileID  Creates and Validates ID to be unique
func ValidateProjectileID(size int, projectiles map[int]*entity.Projectile) int {
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
