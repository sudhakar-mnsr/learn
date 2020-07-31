package main

import (
   "flag"
   "fmt"
   "io"
   "log"
   "net"
   "strings"
   curr "currency/lib0
)

var currencies = curr.Load("../../../data.csv")

// Focus:
// This version of currency server focuses on implementing streaming
// strategy when receiving data from client to avoid dropping data
// when request is larger than internal buffer. This relies on the fact
// that net.Conn implements io.Reader

func main() {
	var addr string
	var network string
	flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
	flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
	flag.Parse()

	// validate supported network protocols
	switch network {
	case "tcp", "tcp4", "tcp6", "unix":
	default:
		log.Fatalln("unsupported network protocol:", network)
	}

ln, err := net.Listen(network, addr)
if err != nil {
   log.Fatal("failed to create listener:", err)
}
defer ln.Close()
log.Println("***** Global Currency Service *****")
log.Printf("Service started: (%s) %s\n", network, addr)
