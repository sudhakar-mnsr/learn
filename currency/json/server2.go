package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	curr "currency/lib"
)

var (
	currencies = curr.Load("../../../data.csv")
)

// This program implements simple currency lookup service over TCP
// or UDS. Clients send curency requests as JSON objects such as 
// {"GET":"USD"}. The request data is then unmarshalled to Go type
// curr.CurrencyRequest("GET":"USD"} using encoding/json
// The search result a []curr.Currency, is marshalled to JSON array
// of objects and send to client.
// IO Streaming:
// This version of server highlights the use of IO streaming when
// using net.Conn to stream data to and from clients
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
