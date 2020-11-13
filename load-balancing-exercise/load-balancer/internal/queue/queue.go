// Package queue enables one-directional communication between frontend and
// backend of the load balancer
package queue

import "github.com/agronskiy/golang-learning/load-balancing-exercise/consts"

// Request is a struct that is sent into the queue
type Request struct {
	Input  int32
	Output chan *Reply
	Err    chan error
}

// Reply is the reply to the request, will be put onto the individual channels
type Reply struct {
	Answer     int32
	WorkerHost string
}

type StopWorkers struct{}

var MainQueue = make(chan *Request)
var StopWorkersChan = make(chan StopWorkers, consts.MaxListeners)
