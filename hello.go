package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:1234/hello"

	// Make a GET request to the server
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return
	}

	// Print the "Hello, World!" message received from the server
	fmt.Println("Response from server:", string(body))
}
