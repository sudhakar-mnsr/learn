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
	// response 48-byte long slice for incoming datagram
	// with time values from the server
	rsp := make([]byte, 48)

	// create an address of type UDPAddr that represents
	// the remote host endpoint
	raddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// setup connection (net.UDPConn) with net.DialUDP()
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("failed while closing connection:", err)
		}
	}()

	fmt.Printf("time from (udp) %s\n", conn.RemoteAddr())

	// Once connection is established, the code pattern
	// is the same as in the other impl.

	// send time request
	if _, err = conn.Write(req); err != nil {
		fmt.Printf("failed to send request: %v\n", err)
		os.Exit(1)
	}

	// block to receive server response
	read, err := conn.Read(rsp)
	if err != nil {
		fmt.Printf("failed to receive response: %v\n", err)
		os.Exit(1)
	}
	//ensure we read 48 bytes back (NTP protocol spec)
	if read != 48 {
		fmt.Println("did not get all expected bytes from server")
		os.Exit(1)
	}
