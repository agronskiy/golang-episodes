package main

import (
	"fmt"
	"time"

	backserver "github.com/agronskiy/golang-learning/load-balancing-exercise/load-balancer/internal/backend-server"
	frontserver "github.com/agronskiy/golang-learning/load-balancing-exercise/load-balancer/internal/frontend-server"
	"github.com/agronskiy/golang-learning/load-balancing-exercise/load-balancer/internal/queue"
)

func main() {

	backserver.InitializeGrpcRegistrationServer()
	frontserver.RunServer()

	for i := 0; i < 10; i++ {
		output := make(chan *queue.Reply)
		err := make(chan error)
		queue.MainQueue <- &queue.Request{Input: 10, Output: output, Err: err}

		reply := <-output
		fmt.Printf("Result: %v from %v\n", reply.Answer, reply.WorkerHost)

		time.Sleep(3 * time.Second)
	}
}
