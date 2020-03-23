package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

// This program implements a simple echo client over TCP or unix domain socket.
// It sends a text content to the server and displays
// the response on the screen.
//
// Usage:
// echoc3 [flags] <text content>
// flags:
//   -e <address-endpoint>
//   -n <network>
func main() {
	var addr string
	var network string
	flag.StringVar(&addr, "e", "localhost:4040", "service address endpoint")
	flag.StringVar(&network, "n", "tcp", "network protocol to use")
	flag.Parse()
	text := flag.Arg(0)
