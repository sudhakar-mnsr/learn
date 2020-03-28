package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	curr "github.com/vladimirvivien/go-networking/currency/lib"
)

var (
	currencies = curr.Load("../data.csv")
)
// This program implements a simple currency lookup service
// over TCP or Unix Data Socket. It loads ISO currency
// information using package curr (see above) and uses a simple
// JSON-encode text-based protocol to exchange data with a client.
//
// Clients send currency search requests as JSON objects
// as {"Get":"<currency name,code,or country"}. The request data is
// then unmarshalled to Go type curr.CurrencyRequest using
// the encoding/json package.
//
// The request is then used to search the list of
// currencies. The search result, a []curr.Currency, is marshalled
// as JSON array of objects and sent to the client.
//
// Focus:
// This version of the code improves on the robustness of the server
// by introducing code to parse the network error and implement connection
// retry logic.  When establishing connection with the client, if that fails
// we can check to see if it is a temporary failure and attempt to retry the
// connection.
//
// Testing:
// Netcat can be used for rudimentary testing.  However, use clientjsonX
// programs functional tests.
//
// Usage: server [options]
// options:
//   -e host endpoint, default ":4040"
//   -n network protocol [tcp,unix], default "tcp"

func main() {
	// setup flags
	var addr string
	var network string
	flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
	flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
	flag.Parse()

	// validate supported network protocols
	switch network {
	case "tcp", "tcp4", "tcp6", "unix":
	default:
		fmt.Println("unsupported network protocol")
		os.Exit(1)
	}

	// create a listener for provided network and host address
	ln, err := net.Listen(network, addr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer ln.Close()
	log.Println("**** Global Currency Service ***")
	log.Printf("Service started: (%s) %s\n", network, addr)

	// delay to sleep when accept fails with a temporary error
	acceptDelay := time.Millisecond * 10
	acceptCount := 0
