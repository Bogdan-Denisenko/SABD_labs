package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	serverCert, err := ioutil.ReadFile("C:\\Users\\mi\\GolandProjects\\SABD_labs\\lab3\\client\\server.pem")
	if err != nil {
		fmt.Println("Error reading server certificate:", err)
		return
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(serverCert)

	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	resp, err := client.Get("https://localhost:8443/hello")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
