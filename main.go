package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	middleware "github.com/ranefattesingh/middleware_demo/middlewares"
)

const ADDRESS string = "localhost:8080"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("The current time is %v", curTime)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", HelloHandler)
	mux.HandleFunc("/v1/time", CurrentTimeHandler)

	// Wrap entire mux in logging middleware
	wrappedMux := middleware.NewResponseHeader(middleware.NewLogger(mux), "x-auth-key", "my-secret-key")

	log.Println(fmt.Sprintf("Server is listening at %s", ADDRESS))
	log.Fatal(http.ListenAndServe(ADDRESS, wrappedMux))
}
