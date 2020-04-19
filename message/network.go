package message

import "sockets/types"

type NetworkMessage struct {
	Data string `json:"data"`
}

//NetworkInput (6 bytes)
type NetworkInput struct {
	IsShooting bool          `json:"IsShooting"`
	Direction  int           `json:"Direction"`
	SequenceID int16         `json:"SequenceID"`
	Rotation   *types.Vector `json:"Rotation"`
	ID         int           `json:"ID"`
}

//UpdatePlayer <- updates player based on input
