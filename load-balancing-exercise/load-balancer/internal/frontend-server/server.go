package srv

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agronskiy/golang-learning/load-balancing-exercise/load-balancer/internal/queue"
)

func doNothing(w http.ResponseWriter, r *http.Request) {}

func handleRequests(w http.ResponseWriter, r *http.Request) {

	log.Println("HandleRequest")
	output := make(chan *queue.Reply)
	err := make(chan error)
	queue.MainQueue <- &queue.Request{Input: 10, Output: output, Err: err}

	reply := <-output
	fmt.Fprintf(w, "Result: %v from %v\n", reply.Answer, reply.WorkerHost)
}

// RunServer starts the server
func RunServer() {
	log.Println("Starting frontend REST server")
	http.HandleFunc("/", handleRequests)
	http.HandleFunc("/favicon.ico", doNothing)

	go func() {
		log.Fatal(http.ListenAndServe(":8090", nil))
	}()

	log.Println("Done")
}
