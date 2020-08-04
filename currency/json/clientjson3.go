package main

import (
	"encoding/json"
	"flag"
	"fmt"
	curr1 "currency/lib"
	"net"
	"os"
	"time"

	curr "currency/lib"
)

const prompt = "currency"

// This program is client implementation for currency service
// request is sent in JSON and response is marshalled in JSON array
// This is over TCP or UDS

// Focus:
// This program uses io streaming, data serilization, client-side
// error handling. Configuration of dialer such as timeout, KeepAlive
// Further code also implements connection retry 

func main() {
var addr string
var network string
flag.StringVar(&addr, "e", "localhost:4040", "service endpoint [ip addr or socket path"])
flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix])
flag.Parse()

// create a dialer to configure its settings instead of using
// the default dialer from net.Dial() function
dialer := &net.Dialer{
   Timeout: time.Second * 300,
   KeepAlive: time.Minute * 5,
}
