package main

import (
	"errors"
	"fmt"

	"github.com/agronskiy/golang-learning/load-balancing-exercise/consts"
)

var (
	freePorts         []string
	downstreamServers map[string]interface{}
)

func initialize() {
	freePorts = make([]string, 0, consts.MaxListeners)
	basePort := 50001
	for i := 0; i < cap(freePorts); i++ {
		freePorts = append(freePorts, fmt.Sprintf("%v", basePort+i))
	}
}

func getNextFreePort() (string, error) {
	if len(freePorts) == 0 {
		return "", errors.New("No free ports")
	}

	res := freePorts[len(freePorts)-1]
	freePorts = freePorts[:len(freePorts)-1]
	return res, nil

}

func main() {
	initialize()
	initializeGrpcRegistrationServer()
}
