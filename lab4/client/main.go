package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	port := 8080
	baseURL := fmt.Sprintf("http://localhost:%d", port)

	sendRequest(baseURL + "/service1")

	sendRequest(baseURL + "/service2")
}

func sendRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("Response from %s: %s\n", url, string(body))
}
