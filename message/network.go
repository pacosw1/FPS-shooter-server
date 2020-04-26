package message

import "sockets/types"

//NetworkMessage t
type NetworkMessage struct {
	Data string `json:"data"`
}

//NetworkInput (6 bytes)
type NetworkInput struct {
	IsShooting bool          `json:"IsShooting"`
	Direction  *types.Point  `json:"Direction"`
	SequenceID uint32        `json:"SequenceID"`
	Rotation   *types.Vector `json:"Rotation"`
	ID         uint32        `json:"ID"`
}

//UpdatePlayer <- updates player based on input
