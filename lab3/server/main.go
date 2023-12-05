package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)

	serverCrt := "C:\\Users\\mi\\GolandProjects\\SABD_labs\\lab3\\server\\server.crt"
	serverKey := "C:\\Users\\mi\\GolandProjects\\SABD_labs\\lab3\\server\\server.key"

	cert, err := tls.LoadX509KeyPair(serverCrt, serverKey)
	if err != nil {
		fmt.Println("Error loading certificate:", err)
		return
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on https://localhost:8443")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
