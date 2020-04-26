package entity

import pb "sockets/protobuf"

//Broadcast s
type Broadcast struct {
	Players     map[uint32]*pb.Player
	Projectiles map[uint32]*pb.Projectile
}

//Zombie t
