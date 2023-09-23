package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	// Create a new gocron scheduler
	s := gocron.NewScheduler(time.UTC)

	// Schedule a job to call the API endpoint every 5 seconds
	_, err := s.Every(5).Seconds().Do(callAPI)
	if err != nil {
		fmt.Println("Error scheduling job:", err)
		return
	}

	// Start the scheduler
	s.StartBlocking()
}

func callAPI() {
	// Define the API endpoint URL
	apiURL := "http://localhost:3000"

	// Send a GET request to the API endpoint
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		fmt.Println("API call successful!")
	} else {
		fmt.Printf("API call failed with status code %d\n", resp.StatusCode)
	}
}
