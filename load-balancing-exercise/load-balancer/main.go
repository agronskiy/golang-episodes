package main

import (
	backserver "github.com/agronskiy/golang-learning/load-balancing-exercise/load-balancer/internal/backend-server"
	frontserver "github.com/agronskiy/golang-learning/load-balancing-exercise/load-balancer/internal/frontend-server"
)

func main() {

	backserver.InitializeGrpcRegistrationServer()
	frontserver.RunServer()

	quit := make(<-chan struct{})
	<-quit
}
