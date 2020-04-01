package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	curr "github.com/vladimirvivien/go-networking/currency/lib0"
)

var currencies = curr.Load("../data.csv")

// This program implements a simple currency lookup service
// over TCP or Unix Data Socket. It loads ISO currency
// information using package lib (see above) and uses a simple
// text-based protocol to interact with the client and send
// the data.
//
// Clients send currency search requests as a textual command in the form:
//
// GET <currency, country, or code>
//
// When the server receives the request, it is parsed and is then used
// to search the list of currencies. The search result is then printed
// line-by-line back to the client.
//
// Focus:
// This version of the currency server focuses on implementing a streaming
// strategy when receiving data from client to avoid dropping data when the
// request is larger than the internal buffer. This relies on the fact that
// net.Conn implements io.Reader which allows the code to stream data.
//
// Testing:
// Netcat or telnet can be used to test this server by connecting and
// sending command using the format described above.
//
// Usage: server0 [options]
// options:
//   -e host endpoint, default ":4040"
//   -n network protocol [tcp,unix], default "tcp"

func main() {
	var addr string
	var network string
	flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
	flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
	flag.Parse()

	// validate supported network protocols
	switch network {
	case "tcp", "tcp4", "tcp6", "unix":
	default:
		log.Fatalln("unsupported network protocol:", network)
	}
	// create a listener for provided network and host address
	ln, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal("failed to create listener:", err)
	}
	defer ln.Close()
	log.Println("**** Global Currency Service ***")
	log.Printf("Service started: (%s) %s\n", network, addr)
	// connection-loop - handle incoming requests
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			if err := conn.Close(); err != nil {
				log.Println("failed to close listener:", err)
			}
			continue
		}
		log.Println("Connected to", conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("error closing connection:", err)
		}
	}()

	if _, err := conn.Write([]byte("Connected...\nUsage: GET <currency, country, or code>\n")); err != nil {
		log.Println("error writing:", err)
		return
	}
