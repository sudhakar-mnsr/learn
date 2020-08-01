package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	curr "currency/lib0"
)

var currencies = curr.Load("../../../data.csv")

// This is simple currency lookup service over TCP or Unix Data Socket
// Text based protocol designed to work on top of TCP or UDS
// Focus:
// There is no streaming strategy for read/write operations
// Buffers are read in one shot creating opportunities for missing data

func main() {
	var addr string
	var network string
	flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
	flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
	flag.Parse()
