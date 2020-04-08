package network

import (
	"github.com/gorilla/websocket"
)

//Client t
type Client struct {
	ID   int
	Conn *websocket.Conn
}

//NewClient creates new client struct instance
func NewClient(ID int, c *websocket.Conn) *Client {
	return &Client{
		ID:   ID,
		Conn: c,
	}
}
