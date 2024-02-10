package main

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	// cancel ensures any requests are closed when parent fuction returns
	defer cancel()

	// two channels
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		// mirrorUrl := url

		// go routines function
		go func(mirrorUrl string) {
			start := time.Now()
			// get request first without & then with the context needed to cancel slower go routines
			// _, err := http.Get(mirrorUrl + "/README")
			req, err := http.NewRequestWithContext(ctx, "GET", mirrorUrl+"/README", nil)
			if err != nil {
				return
			}
			// make the request explicitly now that context is in use
			r, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf("this request was cancled: %s", err)
				return
			}
			// close the response body
			defer r.Body.Close()

			latency := time.Since(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorUrl
				latencyChan <- latency
				// cancel routines spawned, we don't need them anymore, they're too slow
				cancel()
			}
		}(url)
	}

	return response{<-urlChan, <-latencyChan}

}
