package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/agronskiy/golang-learning/load-balancing-exercise/consts"
	pb "github.com/agronskiy/golang-learning/load-balancing-exercise/grpc"
)

var (
	freePorts         []string
	downstreamServers map[string]interface{}
)

func init() {
	freePorts = make([]string, 0, consts.MaxListeners)
	basePort := 50001
	for i := 0; i < cap(freePorts); i++ {
		freePorts = append(freePorts, fmt.Sprintf("%v", basePort+i))
	}
}

type registrationServer struct {
	pb.UnimplementedRegistrarServer
}

func getNextFreePort() (string, error) {
	if len(freePorts) == 0 {
		return "", errors.New("No free ports")
	}

	res := freePorts[len(freePorts)-1]
	freePorts = freePorts[:len(freePorts)-1]
	return res, nil

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

// InitializeGrpcRegistrationServer initializes the gRPC service
func InitializeGrpcRegistrationServer() {
	log.Printf(fmt.Sprintf("localhost:%s", consts.BackendPort))
	lis, err := net.Listen(consts.BackendProtocol, fmt.Sprintf("localhost:%s", consts.BackendPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	log.Println("Initializing gRPC backend registration server NEW...")
	pb.RegisterRegistrarServer(grpcServer, &registrationServer{})
	log.Println("Starting gRPC backend registration server...")
	grpcServer.Serve(lis)
	log.Println("Done")
}
