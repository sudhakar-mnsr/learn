package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// This program is a simple Network Time Protocol server over UDP.
// The implementation uses a UDPConn and ListenUDP to manage requests.
// The server returns the number of seconds since 1900 up to the
// current time.

// Again this is a simple server, it dies after sending the response.

// Usage:
// ntps -e <host address endpoint>
func main() {
	var host string
	flag.StringVar(&host, "e", ":1123", "server address")
	flag.Parse()

	// Create a UDP host addres
	addr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
