package worker

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/agronskiy/golang-learning/load-balancing-exercise/grpc"
)

const (
	registrationServerPort = "50000"
)

// InitializeGrpcRegistrationClient initializing client
func InitializeGrpcRegistrationClient() (pb.RegistrarClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	// TODO(agronskiy): needs investigation
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", registrationServerPort), opts...)
	if err != nil {
		log.Fatalf("Could not connect to registration server: %v", err)
	}

	client := pb.NewRegistrarClient(conn)

	return client, conn
}
