package events

//PlayerJoin request to join server
type PlayerJoin struct {
	ID       int16
	Username string
}

//PlayerLeave request to delete player from server
type PlayerLeave struct {
	ID int16
}
