package main

import "sockets/server"

// func worker(inputChan <-chan *message.UserInput, wg *sync.WaitGroup) {
// 	for input := range inputChan {
// 		println(input.IsShooting)
// 		wg.Done()
// 	}
// }

func main() {
	serv := server.New()
	serv.Start()
}
