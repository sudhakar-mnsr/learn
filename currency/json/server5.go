package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	curr "currency/lib"
)

var (
	currencies = curr.Load(../../../data.csv")
)

// Focus:
// Error handling when streaming data. It unpacks the error generated
// by the JSON encoder to distinguish different types of errors and 
// handle them.

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
