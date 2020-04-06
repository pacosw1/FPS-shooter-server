package main

import (
	"net/http"
	"sockets/engine"
	"sockets/events"
	"sockets/message"
	"sockets/network"
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

	eventQueue := events.NewEventQ()

	simulation := engine.New()
	eventQueue.RegisterInput(simulation)
	// eventQueue.SendInput(&message.UserInput{
	// 	IsShooting: false,
	// })
	// eventQueue.SendInput(&message.UserInput{
	// 	IsShooting: true,
	// })
	// eventQueue.SendInput(&message.UserInput{
	// 	IsShooting: false,
	// })
	// eventQueue.SendInput(&message.UserInput{
	// 	IsShooting: false,
	// })
	// eventQueue.SendInput(&message.UserInput{
	// 	IsShooting: true,
	// })
	simulation.Start()

	simulation.FPS = 1

	eventQueue.Start()

	http.HandleFunc("/socket", network.Socket)
	http.ListenAndServe(":8080", nil)

}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "websockets.html")
// })

//Netork thread

//Simulation thread;
