package worker

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/agronskiy/golang-learning/load-balancing-exercise/grpc"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

type mockServer struct {
	pb.UnimplementedRegistrarServer
}

func (s *mockServer) RequestWorkerRegistration(
	ctx context.Context,
	r *pb.RegistrationRequest,
) (*pb.RegistrationReply, error) {

	reply := &pb.RegistrationReply{Ok: true}
	return reply, nil
}

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterRegistrarServer(s, &mockServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestRequestRegistration(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewRegistrarClient(conn)
	resp, err := client.RequestWorkerRegistration(ctx, &pb.RegistrationRequest{Port: "0000"})
	if err != nil {
		t.Fatalf("RequestWorkerRegistration call failed: %v", err)
	}

	if resp.Ok != true {
		t.Errorf("RequestWorkerRegistration() returned `false` when otherwise requested.")
	}
}
