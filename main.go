package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request from %s", GetIP(r))
	time.Sleep(time.Duration(getDelay()) * time.Millisecond)

	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
	fmt.Fprintf(w, "Hopefully I'm containerized.\n")
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprint(w, "I may not get my hostname.\n")
	} else {
		fmt.Fprintf(w, "Hostname: %s\n", hostname)
	}
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func getDelay() int64 {
	delayS := os.Getenv("DELAY")
	delay, err := strconv.Atoi(delayS)
	if err != nil && delayS != "" {
		log.Println("Delay not proper set.")
		return 0
	}
	return int64(delay)
}
