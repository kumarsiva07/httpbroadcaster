package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"sync"
	"time"
)

func main() {
	h := &baseHandle{}
	http.Handle("/", h)

	// handle all requests to your multiserver using the proxy
	log.Fatal(http.ListenAndServe(":8888", h))

}

var hostTarget = []string{
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8090",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
	"http://localhost:8091",
}

type baseHandle struct{}

func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// create a WaitGroup
	wg := new(sync.WaitGroup)

	wg.Add(len(hostTarget))
	for _, host := range hostTarget {
		go func() {
			process(r, host)
			wg.Done()
		}()
	}
	wg.Wait()
	w.Write([]byte("OK"))
}

func process(req *http.Request, host string) {

	//transport := &http.Transport{}
	//client := &http.Client{Transport: transport}
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	dump, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(dump))
	request, err := http.NewRequest(req.Method, host+req.URL.String(), bytes.NewReader(dump))

	response, err := client.Do(request)
	if err == nil {
		readResponse(response)
	}
	return
}

func readResponse(response *http.Response) {
	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			break
		}
		fmt.Print(string(buf[0:n]))
	}
}
