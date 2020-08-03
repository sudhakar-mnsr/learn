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

	curr "currency/lib"
)

var (
	currencies = curr.Load("../../../data.csv")
)

// This program implements simple currency lookup service over TCP or UDS
// The request data is unmarshalled to curr.CurrencyRequest
// The search result is marshalled as JSON array and sent to client.
// Focus:
// This code is more roboust in parsing network error and implement retry
// On temporary failures the server attempt to retry connecting.

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

ln, err := net.Listen(network, addr)
if err != nil {
   log.Println(err)
   os.Exit(1)
}
defer ln.Close()
log.Println("***** Global Currency Service *****")
log.Printf("Service started: (%s) %s\n", network, addr)

// delay to sleep when accept fails with a temporary error
acceptDelay := time.Millisecond * 10
acceptCount := 0
