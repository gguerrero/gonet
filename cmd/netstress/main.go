package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var uriFlag = flag.String("url", "", "URL where the stress test is going to point, using GET")
var maxRoutinesFlag = flag.Int("max", 10, "max concurrent requests run at a time")

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	var wg sync.WaitGroup

	flag.Parse()

	for i := 0; i < *maxRoutinesFlag; i++ {
		wg.Add(1)
		go timeRequest(*uriFlag, i+1, &wg)
	}

	wg.Wait()
}

func timeRequest(uri string, n int, wg *sync.WaitGroup) {
	defer wg.Done()

	reqStartTs := time.Now()

	r, err := http.Get(uri)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	reqEndTs := time.Now()

	reqDuration := reqEndTs.Sub(reqStartTs)
	log.Printf(
		"Request #%d - Duration %d µs (%d ms) - Status %d\n",
		n, reqDuration.Microseconds(), reqDuration.Milliseconds(), r.StatusCode,
	)
}
