package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Req host : " + req.Host)
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL)
		h.ServeHTTP(w, r)
	})
}

func createServer(name string, port int) *http.Server {

	// create `ServerMux`
	mux := http.NewServeMux()

	// create a default route handler
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)

	// create new multiserver
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port), // :{port}
		Handler: mux,
	}

	// return new multiserver (pointer)
	return &server
}

func main() {

	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two goroutines to `wg` WaitGroup
	wg.Add(2)

	// goroutine to launch a multiserver on port 9000
	go func() {
		server := createServer("ONE", 8090)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	// goroutine to launch a multiserver on port 9001
	go func() {
		server := createServer("TWO", 8091)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	// wait until WaitGroup is done
	wg.Wait()

}
