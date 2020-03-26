package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	curr "github.com/vladimirvivien/go-networking/tcp/curlib"
)

var (
	currencies = curr.Load("./data.csv")
)

// This program implements a simple currency lookup service
// over TCP or Unix Data Socket. It loads ISO currency
// information using package curlib (see above) and makes
// and serves it using JSON-enocoded data.
//
// Clients send currency search requests as JSON objects such
// as {"Get":"USD"}. The request data is then unmarshalled to Go
// type curr.CurrencyRequest{Get:"USD"} using the encoding/json
// package.
//
// The request is then used to search the list of
// currencies. The search result, a []curr.Currency, is marshalled
// to JSON array of objects and send to the client.
//
// IO Streaming:
// This version of the server highlights the use of IO streaming
// when using net.Conn to stream data to and from clients.
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
