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
// search requests JSON objects unmarshalled to curr.CurrencyRequest
// search results curr.Currency marshalled to JSON array and sent back
// Focus:
// This version improve robustness by introducing configuration for read
// and write timeout values. This ensures that client cannot hold a 
// connection hostage by taking long time to send or receive data

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
