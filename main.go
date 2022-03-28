package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	host = "localhost"
	port = "8090"
)

func main() {
	listenURL := fmt.Sprintf("%s:%s", host, port)

	http.HandleFunc("/", ResponseHandler)

	fmt.Printf("Server started at port %s\n", listenURL)
	log.Fatal(http.ListenAndServe(listenURL, nil))
}

// ResponseHandler returns the default HTML code response for this container
func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		`
		<h1>Hello, there</h1>
		I'm running on: <b>%s</b>
		`, getHostname())
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return err.Error()
	}
	return hostname
}
