package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	defaultMaxRoutines = 1000
)

type options struct {
	uri         string
	maxRoutines int
}

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	var opts = parseOpts()
	var wg sync.WaitGroup

	for i := 0; i < opts.maxRoutines; i++ {
		wg.Add(1)
		go timeRequest(opts.uri, i+1, &wg)
	}

	wg.Wait()
}

func parseOpts() *options {
	args := os.Args[1:]
	opts := &options{
		uri:         args[0],
		maxRoutines: defaultMaxRoutines,
	}

	maxRoutines, _ := strconv.Atoi(args[1])
	if maxRoutines > 0 {
		opts.maxRoutines = maxRoutines
	}

	return opts
}

func timeRequest(uri string, n int, wg *sync.WaitGroup) {
	defer wg.Done()

	reqStartTs := time.Now()

	r, err := http.Get(uri)
	if err != nil {
		log.Println(err)
	}

	reqEndTs := time.Now()

	reqDuration := reqEndTs.Sub(reqStartTs)
	log.Printf(
		"Request #%d - Duration %d µs (%d ms) - Status %d\n",
		n, reqDuration.Microseconds(), reqDuration.Milliseconds(), r.StatusCode,
	)
}
