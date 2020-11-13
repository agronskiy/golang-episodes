package server

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/agronskiy/golang-learning/load-balancing-exercise/grpc"
	"github.com/agronskiy/golang-learning/load-balancing-exercise/load-balancer/internal/queue"
)

type workerGrpcInfo struct {
	client pb.WorkerClient
	conn   *grpc.ClientConn
}

func workerGoroutine(
	host string,
	quit <-chan queue.StopWorkers,
	inputQueue <-chan *queue.Request,
) {
	workerGrpc := initializeWorkerGrpcClient(host)
	defer workerGrpc.conn.Close()

	var req *queue.Request
	for {
		select {
		case req = <-inputQueue:
			reply, err := workerGrpc.client.PerformJob(
				context.Background(),
				&pb.JobRequest{X: req.Input},
			)
			if err != nil {
				req.Err <- err
				continue
			}
			req.Output <- &queue.Reply{Answer: reply.Result, WorkerHost: host}
		case <-quit:
			log.Printf("Closing connection to host: %v", host)
			return
		}
	}
}

func initializeWorkerGrpcClient(host string) *workerGrpcInfo {
	var opts []grpc.DialOption
	// TODO(agronskiy): needs investigation
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Fatalf("Could not connect to worker's gRPC server: %v", err)
	}

	client := pb.NewWorkerClient(conn)

	return &workerGrpcInfo{client: client, conn: conn}
}
