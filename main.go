package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "Hopefully I'm containerized.")
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprint(w, "I may not get my hostname.")
	} else {
		fmt.Fprintf(w, "Hostname: %s", hostname)
	}
}
