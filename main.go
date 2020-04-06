package main

import (
	"net/http"
	"sockets/events"
	"sockets/message"
	"sockets/network"
	"sockets/state"
	"sync"
)

var wg sync.WaitGroup

func worker(inputChan <-chan *message.UserInput, wg *sync.WaitGroup) {
	for input := range inputChan {
		println(input.IsShooting)
		wg.Done()
	}
}

func main() {

	gameState := state.New()

	eventQueue := events.NewEventQ()

	eventQueue.RegisterConnect(gameState)
	eventQueue.RegisterDisconnect(gameState)

	go eventQueue.Start()

	eventQueue.FireConnect(message.ConnectMessage("pacosw1", 1))

	eventQueue.FireDisconnect(message.DisconnectMessage(1))

	http.HandleFunc("/socket", network.Socket)
	http.ListenAndServe(":8080", nil)

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
