package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	host = ""
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
	fmt.Printf("Request from %s(%s)\n", r.Host, r.RemoteAddr)
	fmt.Fprintf(w,
		`
<h1>Hello, there</h1>
<div>
I'm the <i>avillega-hello-world:</i><b>v2.0</b>
</div>
<div>
I'm running on:
</div>
<div>
Version: <b>%s</b>
</div>
<div>
Process ID: <b>%d</b>
</div>
<div>
User ID: <b>%d</b>
</div>
<div>
Group ID: <b>%d</b>
</div>

`, getHostname(), getPID(), getUID(), getGID())
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return err.Error()
	}
	return hostname
}

func getPID() int {
	return os.Getpid()
}

func getUID() int {
	return os.Getuid()
}

func getGID() int {
	return os.Getgid()
}
