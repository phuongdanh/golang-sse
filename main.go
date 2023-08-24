package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	eventChan := make(chan string) // Channel to send events to clients

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		for {
			eventData := <-eventChan
			fmt.Fprintf(w, "data: %s\n\n", eventData)
			w.(http.Flusher).Flush()
		}
	})

	http.HandleFunc("/send-event", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		eventData := fmt.Sprintf("Button clicked at %s", time.Now().Format("2006-01-02 15:04:05"))
		eventChan <- eventData
		w.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start the server:", err)
		return
	}
}
