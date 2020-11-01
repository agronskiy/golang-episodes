package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/agronskiy/golang-learning/load-balancing-exercise/grpc"
)

const (
	registrationServerPort = "50000"
)

func main() {
	client, conn := initializeGrpcRegistrationClient()
	defer conn.Close()

	reply, err := client.RequestWorkerRegistration(
		context.Background(),
		&pb.RegistrationRequest{Host: "localhost"})
	if err != nil {
		log.Fatalf("Could not request port: %v", err)
	}
	fmt.Printf("Reply: %s", reply.Port)
}
