package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"learning-go/ch11/curr1"
	"net"
	"os"
	"time"

	curr "currency/lib"
)

const prompt = "currency"

// This program is a client implementation for the currency service
// program. It sends JSON requests and receives JSON-encoded array
// over TCP or UDS
// FOCUS:
// IO streaming
// data serialization
// client-side error handling
// IMPORTANT: configure dialer to setup settings such as timeout and 
// KeepAlive values. Further the code implements a simple connection-retry
// strategy when connecting

func main() {
var addr string
var network string
flag.StringVar(&addr, "e", "localhost:4040", "service endpoint [ip addr or socket path]")
flag.StringVar(&network, "n", "tcp", "network protocol [tcp, unix]")
flag.Parse()

// create a dialer to configure its settings instead of using the default 
// dialer from net.Dial() function
dialer := &net.Dialer{
   Timeout: time.Second * 300,
   KeepAlive: time.Minute * 5,
}
