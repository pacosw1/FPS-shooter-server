package utils

import (
	"sockets/entity"

	"google.golang.org/protobuf/proto"
)

//MarshalMessage convert message
func MarshalMessage(message proto.Message) *[]byte {
	bytes, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}

	return &bytes
}

//CopyState copies state

//CopyPlayers copy player state
func CopyPlayers(players map[int]*entity.Player) map[int]*entity.Player {
	copy := make(map[int]*entity.Player)

	for key, value := range players {
		copy[key] = value
	}

	return copy

}

//CopyProjectiles cp
func CopyProjectiles(projectiles map[int]*entity.Projectile) map[int]*entity.Projectile {
	copy := make(map[int]*entity.Projectile)

	for key, value := range projectiles {
		copy[key] = value
	}
	return copy
}
