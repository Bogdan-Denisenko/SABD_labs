package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

const customHeader = "X-Custom-Header"

type Route struct {
	Prefix     string
	BackendURL string
}

var routes = []Route{
	{Prefix: "/service1", BackendURL: "http://localhost:8081"},
	{Prefix: "/service2", BackendURL: "http://localhost:8082"},
}

func NewReverseProxy(targetURL *url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.Director = func(req *http.Request) {
		req.Header.Add(customHeader, "test")
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = singleJoiningSlash(targetURL.Path, req.URL.Path)
	}

	return proxy
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func GatewayHandler(routes []Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, route := range routes {
			if strings.HasPrefix(r.URL.Path, route.Prefix) {
				targetURL, _ := url.Parse(route.BackendURL)
				proxy := NewReverseProxy(targetURL)
				proxy.ServeHTTP(w, r)
				return
			}
		}

		http.NotFound(w, r)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/{path:.*}", GatewayHandler(routes))

	port := 8080
	fmt.Printf("Gateway is listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
