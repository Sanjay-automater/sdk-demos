package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("Please set the API_KEY environment variable.")
		os.Exit(1)
	}

	endpoint := "https://api.example.com/v1/demo" // placeholder

	payload := map[string]string{"message": "hello from go demo"}
	b, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Failed to build request:", err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.StatusCode)
	// For brevity we omit reading the whole body here
}
