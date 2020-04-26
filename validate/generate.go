package validate

import (
	"math/rand"
	"time"
)

// PlayerID Creates and Validates ID to be unique

//GenerateID a unique ID beteen 0 and 100
func GenerateID(size int) uint32 {
	timestamp := time.Now().UnixNano()
	rand.Seed(timestamp)
	id := uint32(rand.Intn(size))
	return id
}
