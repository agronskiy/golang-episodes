package srv

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sample response")
}

// RunServer starts the server
func RunServer() {
	log.Println("Starting frontend REST server")
	http.HandleFunc("/", handleRequests)
	go func() {
		log.Fatal(http.ListenAndServe(":8090", nil))
	}()

	log.Println("Done")
}
