package main

import (
	"net/http"
	"sockets/engine"
	"sockets/events"
	"sockets/network"
	"sockets/object"
	"sync"
)

var wg sync.WaitGroup

func worker(inputChan <-chan *object.UserInput, wg *sync.WaitGroup) {
	for input := range inputChan {
		println(input.IsShooting)
		wg.Done()
	}
}

func main() {

	eventQueue := events.New()

	simulation := engine.New()
	eventQueue.RegisterInput(simulation)
	eventQueue.SendInput(&object.UserInput{
		IsShooting: false,
	})
	eventQueue.SendInput(&object.UserInput{
		IsShooting: true,
	})
	eventQueue.SendInput(&object.UserInput{
		IsShooting: false,
	})
	eventQueue.SendInput(&object.UserInput{
		IsShooting: false,
	})
	eventQueue.SendInput(&object.UserInput{
		IsShooting: true,
	})
	engine.Start(simulation)

	events.Start(eventQueue)

	http.HandleFunc("/socket", network.Socket)
	http.ListenAndServe(":8080", nil)

}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "websockets.html")
// })

//Netork thread

//Simulation thread;
