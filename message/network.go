package message

import "sockets/types"

type NetworkMessage struct {
	Data string `json:"data"`
}
type NetworkInput struct {
	IsShooting bool            `json:"IsShooting"`
	Direction  *types.Position `json:"Direction"`
	SequenceID int16           `json:"SequenceID"`
	Aim        *types.Position `json:"Aim"`
	ID         int             `json:"ID"`
}
