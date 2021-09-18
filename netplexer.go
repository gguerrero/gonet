package main

import (
	"log"
	"net"
	"os"

	"github.com/gguerrero/gonet/netplexer"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	server := netplexer.Server{
		Address: net.ParseIP("0.0.0.0"),
		Port:    8000,
	}

	server.Serve()
}
