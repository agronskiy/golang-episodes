package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"

	"github.com/agronskiy/golang-learning/load-balancing-exercise/consts"
	pb "github.com/agronskiy/golang-learning/load-balancing-exercise/grpc"
	"github.com/agronskiy/golang-learning/load-balancing-exercise/load-balancer/internal/queue"
)

type registrationServer struct {
	pb.UnimplementedRegistrarServer
}

func (*registrationServer) RequestWorkerRegistration(
	ctx context.Context,
	r *pb.RegistrationRequest,
) (*pb.RegistrationReply, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("RegistrationServer: Could not determine the host of the worker")
	}
	workerHost := p.Addr.(*net.TCPAddr).IP.String() + ":" + r.Port
	log.Printf("Worker running on host: %v", workerHost)

	// This starts the worker client which reads from the input channel
	go func() {
		workerGoroutine(workerHost, queue.StopWorkersChan, queue.MainQueue)
	}()

	reply := &pb.RegistrationReply{Ok: true}
	return reply, nil
}

// InitializeGrpcRegistrationServer initializes the gRPC service
func InitializeGrpcRegistrationServer() {
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

	go func() {
		grpcServer.Serve(lis)
	}()

	log.Println("Done")
}
