package main

import entity "sockets/create"

func main() {

	println("hello world")
	playerList := entity.PlayerList()
	entity.AddPlayer(1, 1, playerList)
	entity.AddPlayer(1, 1, playerList)
	entity.AddPlayer(1, 1, playerList)
	entity.AddPlayer(1, 1, playerList)

	println(len(playerList))

}

//Netork thread

//Simulation thread;
