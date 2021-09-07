package main

import (
	"fmt"
	"net/http"
	"sync"
)

func createHTTPServer(name string, port int) *http.Server {
	// Create custom multiplexer
	mux := http.NewServeMux()

	// Create default router handler
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello: "+name)
	})

	// Log server
	fmt.Println("Running HTTP Server on port", port)

	// Create new http server
	return &http.Server{Addr: fmt.Sprintf(":%v", port), Handler: mux}
}

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		server := createHTTPServer("HTTP SERVER 1", 5000)
		server.ListenAndServe()
		wg.Done()
	}()

	go func() {
		server := createHTTPServer("HTTP SERVER 2", 5001)
		server.ListenAndServe()
		wg.Done()
	}()

	wg.Wait()
}
