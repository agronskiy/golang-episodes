package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/agronskiy/golang-learning/load-balancing-exercise/consts"

	"google.golang.org/grpc"

	pb "github.com/agronskiy/golang-learning/load-balancing-exercise/grpc"
)

type registrationServer struct {
	pb.UnimplementedRegistrarServer
}

func (s *registrationServer) RequestWorkerRegistration(
	ctx context.Context,
	r *pb.RegistrationRequest,
) (*pb.RegistrationReply, error) {
	log.Printf("Request from host: %v", r.Host)

	port, err := getNextFreePort()
	if err != nil {
		log.Printf("Error: %s", err)
		reply := &pb.RegistrationReply{Ok: false}
		return reply, err
	}

	reply := &pb.RegistrationReply{Ok: true, Port: port}
	return reply, nil
}

func initializeGrpcRegistrationServer() {
	log.Printf(fmt.Sprintf("localhost:%s", consts.BackendPort))
	lis, err := net.Listen(consts.BackendProtocol, fmt.Sprintf("localhost:%s", consts.BackendPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	log.Println("Initializing gRPC backend registration server...")
	pb.RegisterRegistrarServer(grpcServer, &registrationServer{})
	log.Println("Starting gRPC backend registration server...")
	grpcServer.Serve(lis)
	log.Println("Done")
}
