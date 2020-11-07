package main

import (
	srv "github.com/agronskiy/golang-learning/load-balancing-exercise/load-balancer/internal/backend-server"
)

func main() {
	srv.InitializeGrpcRegistrationServer()
}
