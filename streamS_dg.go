package main

import (
"flag"
"fmt"
"net"
"os"
)

// Simple echo server over TCP or Unix Domain Socket (streaming)
func main() {
var addr string
var network string
flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
flag.StringVar(&network, "n", "tcp", "network protocol [tcp, unix]")
flag.Parse()

switch network {
case "tcp", "tcp4", "tcp6", "unix":
default:
   fmt.Println(unsupported network protocol")
   os.Exit(1)
}
