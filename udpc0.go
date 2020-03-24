package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// This program implements a trivial NTP client over UDP.
// It uses NTP version 3 data packet format which is a
// 48-byte long datagram for both request and response.
// Usage:
// ntpc -e <host endpoint>
func main() {
	var host string
	flag.StringVar(&host, "e", "us.pool.ntp.org:123", "NTP host")
	flag.Parse()

	// req datagram is a 48-byte long slice
	// that is used for sending time request to the server
	req := make([]byte, 48)

	// req is initialized with 0x1B or 0001 1011 which is
	// a request setting for time server.
	// See spec at ntp.org
	req[0] = 0x1B
