package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/agronskiy/golang-learning/load-balancing-exercise/grpc"
	"github.com/agronskiy/golang-learning/load-balancing-exercise/worker/internal/worker"
)

func registerWorker(port string) {
	client, conn := worker.InitializeGrpcRegistrationClient()
	defer conn.Close()

	reply, err := client.RequestWorkerRegistration(
		context.Background(),
		&pb.RegistrationRequest{Port: port})
	if err != nil {
		log.Fatalf("Could not request port: %v", err)
	}
	fmt.Println(fmt.Sprintf("Reply: %v", reply.Ok))
}

func main() {
	port := worker.RunGrpcWorkerServer()
	registerWorker(port)
	for {
	}
}
