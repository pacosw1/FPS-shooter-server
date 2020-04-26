package utils

import (
	"sockets/entity"
	pb "sockets/protobuf"
	"sockets/state"
)

//CopyState copies state
func CopyState(s *state.GameState) *entity.Broadcast {

	ogPlayers := s.Players
	ogProject := s.Projectiles

	players := make(map[uint32]*pb.Player)
	projectiles := make(map[uint32]*pb.Projectile)

	for key, value := range ogPlayers {
		players[key] = value.ToProto()
	}

	for key, value := range ogProject {
		projectiles[key] = value.ToProto()
	}

	return &entity.Broadcast{
		Players:     players,
		Projectiles: projectiles,
	}

}

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
