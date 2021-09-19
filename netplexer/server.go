package netplexer

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/gguerrero/gonet/netplexer/request"
	"github.com/gguerrero/gonet/netplexer/response"
)

// port defines the TCP port where the server is running.
type port uint16

// Server defines the IP and the port where the TCP server will run
type Server struct {
	Address net.IP
	Port    port
}

const (
	network = "tcp"
	// TCP requests default timeout.
	timeout = time.Second * 5
)

// The server name as IP and Port
func (s *Server) String() string {
	return fmt.Sprintf("%s:%d", s.Address, s.Port)
}

// Listen, accept and handles TCP connections.
// HTTP protocol request are expected as input.
func (s *Server) Serve() {
	log.Printf("Netplexer listening at %s\n", s.String())
	li, err := net.Listen(network, s.String())
	if err != nil {
		log.Fatal(err)
	}
	defer closeListener(li)
	handleOSInterrupt(li)

	acceptAndHandle(li)
}

func acceptAndHandle(li net.Listener) {
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer closeConnection(conn)

	reqStartTs := time.Now()

	conn.SetDeadline(time.Now().Add(timeout))

	req, err := request.NewHttpRequest(conn)
	if err != nil {
		log.Fatal(err)
	}

	response.NewhttpResponseWriter(conn).Write(req)

	reqEndTs := time.Now()
	reqDuration := reqEndTs.Sub(reqStartTs)
	log.Printf("Response time %d µs (%d ms) 🚀", reqDuration.Microseconds(), reqDuration.Milliseconds())
}

func closeConnection(conn net.Conn) {
	log.Println("Connection closed on", conn.RemoteAddr())
	conn.Close()
}

func closeListener(li net.Listener) {
	log.Println("Listener closed on", li.Addr())
	li.Close()
}

func handleOSInterrupt(li net.Listener) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig.String() == "interrupt" {
				log.Println("... ^C captured, stopping server!")
				closeListener(li)
				os.Exit(0)
			}
		}
	}()
}
