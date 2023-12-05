package main

import (
	"fmt"
	"net/http"
)

const customHeader = "X-Custom-Header"

func main() {
	port := 8082
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if value := r.Header.Get(customHeader); value != "" {
			fmt.Fprintf(w, "Service 2 - Custom Header Value: %s", value)
		} else {
			fmt.Fprint(w, "Service 2 - No Custom Header")
		}
	})

	fmt.Printf("Service 2 is listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
