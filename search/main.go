package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Check if a search term was provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a search term")
		os.Exit(1)
	}

	// Combine all arguments into the search query
	searchQuery := strings.Join(os.Args[1:], " ")

	// URL encode the search query
	encodedQuery := url.QueryEscape(searchQuery)

	// Create the full Google search URL
	searchURL := "https://www.google.com/search?q=" + encodedQuery

	// Fetch the HTML content
	html, err := fetchHTML(searchURL)
	if err != nil {
		fmt.Printf("Error fetching search results: %v\n", err)
		os.Exit(1)
	}

	// Print the HTML to terminal
	fmt.Println(html)
}

// fetchHTML fetches the HTML content from the given URL
func fetchHTML(url string) (string, error) {
	// Create an HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set a User-Agent to avoid being blocked
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
