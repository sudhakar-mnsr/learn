package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"

	curr "currency/lib"
)

const prompt = "currency"

// It sends JSON-encoded requests and receives JSON array of currency
// information directly over TCP or UDS
// Focus:
// This version of client program highlights the use of IO streaming
// data serilization and client-side error handling

func main() {
var addr string
var network string
flag.StringVar(&addr, "e", "localhost:4040", "service endpoint [ip addr or socket path]")
flag.StringVar(&network, "n", "tcp", "network protocol [tcp, unix]")
flag.Parse()

