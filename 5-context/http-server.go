package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe(":8082", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("starting request")
	defer log.Println("request has finished")
	select {
	case <-time.After(5 * time.Second):
		log.Println("request has processed")
		w.Write([]byte("request has processed successfully"))
	case <-ctx.Done():
		log.Println("request has canceled")
		http.Error(w, "request has been canceled", http.StatusRequestTimeout)
	}
}
