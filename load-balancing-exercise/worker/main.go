package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/agronskiy/golang-learning/load-balancing-exercise/grpc"
	"github.com/agronskiy/golang-learning/load-balancing-exercise/worker/internal/worker"
)

func runClient(c chan int) {
	client, conn := worker.InitializeGrpcRegistrationClient()
	defer conn.Close()

	reply, err := client.RequestWorkerRegistration(
		context.Background(),
		&pb.RegistrationRequest{Host: "localhost"})
	if err != nil {
		log.Fatalf("Could not request port: %v", err)
	}
	fmt.Printf("Reply: %s", reply.Port)

	c <- 0
}

func main() {
	c := make(chan int)
	go runClient(c)

	// Wait until
	<-c
}
