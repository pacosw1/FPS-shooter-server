package main

import (
	"sockets/events"
	"sockets/network"
	"sockets/state"
)

// var wg sync.WaitGroup

// func worker(inputChan <-chan *message.UserInput, wg *sync.WaitGroup) {
// 	for input := range inputChan {
// 		println(input.IsShooting)
// 		wg.Done()
// 	}
// }

func main() {
	eventQueue := events.NewEventQ()
	gameState := state.New(eventQueue)

	net := network.New(eventQueue, gameState)

	eventQueue.RegisterConnect(gameState)
	eventQueue.RegisterBroadcast(net)
	eventQueue.RegisterDisconnect(gameState)
	eventQueue.Start()
	gameState.Start()
	net.Start()

	// eventQueue.FireConnect(message.ConnectMessage("pacosw1", 1))

	// eventQueue.FireDisconnect(message.DisconnectMessage(1))

}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "websockets.html")
// })

//Netork thread

//Simulation thread;

// eventQueue := events.NewEventQ()

// simulate := simulation.NewEngine()
// eventQueue.RegisterInput(simulate)
// // eventQueue.SendInput(&message.UserInput{
// // 	IsShooting: false,
// // })
// // eventQueue.SendInput(&message.UserInput{
// // 	IsShooting: true,
// // })
// // eventQueue.SendInput(&message.UserInput{
// // 	IsShooting: false,
// // })
// // eventQueue.SendInput(&message.UserInput{
// // 	IsShooting: false,
// // })
// // eventQueue.SendInput(&message.UserInput{
// // 	IsShooting: true,
// // })
// simulate.Start()

// simulate.FPS = 1

// eventQueue.Start()
