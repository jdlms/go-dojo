package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

func main() {

	http.HandleFunc("/fastest", func(w http.ResponseWriter, r *http.Request) {
		response := findFastest(MirrorList[:])
		respJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	})

	port := ":8000"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(server.ListenAndServe())

}

func findFastest(urls []string) response {

	// two channels
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		mirrorUrl := url

		// go routines function
		go func() {
			start := time.Now()

			_, err := http.Get(mirrorUrl + "/README")

			if err != nil {
				return
			}
			// make the request explicitly now that context is in use

			latency := time.Since(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorUrl
				latencyChan <- latency
			}
		}()
	}

	return response{<-urlChan, <-latencyChan}

}
