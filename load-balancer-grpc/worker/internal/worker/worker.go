package worker

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/agronskiy/golang-episodes/load-balancer-grpc/consts"
	pb "github.com/agronskiy/golang-episodes/load-balancer-grpc/grpc"
)

// InitializeGrpcRegistrationClient initializing client
func InitializeGrpcRegistrationClient() (pb.RegistrarClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	// TODO(agronskiy): needs investigation
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", consts.BackendPort), opts...)
	if err != nil {
		log.Fatalf("Could not connect to registration server: %v", err)
	}

	client := pb.NewRegistrarClient(conn)

	return client, conn
}

type workerGrpcServer struct {
	pb.UnimplementedWorkerServer
}

func (s *workerGrpcServer) PerformJob(
	ctx context.Context,
	r *pb.JobRequest,
) (*pb.JobReply, error) {
	log.Printf("Requested job, input: %v", r.X)

	return &pb.JobReply{Result: -r.X}, nil
}

// RunGrpcWorkerServer initializes the gRPC service of the worker
func RunGrpcWorkerServer() (port string) {
	lis, err := net.Listen(consts.BackendProtocol, ":0")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	port = fmt.Sprint(lis.Addr().(*net.TCPAddr).Port)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	log.Println("Initializing gRPC worker server...")
	pb.RegisterWorkerServer(grpcServer, &workerGrpcServer{})
	log.Printf("Starting gRPC worker server on port %v...", port)

	go func() {
		grpcServer.Serve(lis)
	}()

	log.Println("Done")

	return port
}
