package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// This program is a simple Network Time Protocol server over
// Unix Domain Socket instead of UDP. The implementation uses
// ListenUnixgram and UnixConn to manage requests.
// The server returns the number of seconds since 1900 up to the
// current time.

// Usage:
// ntps -e <host address endpoint>
func main() {
	var path string
	flag.StringVar(&path, "e", "/tmp/time.sock", "NTP server socket endpoint")
	flag.Parse()

	// Creaets a UnixAddr address
	addr, err := net.ResolveUnixAddr("unixgram", path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// setup connection UnixConn with ListenUnixgram
	conn, err := net.ListenUnixgram("unixgram", addr)
	if err != nil {
		fmt.Println("failed to create socket:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("listening on (unixgram) %s\n", conn.LocalAddr())

	// request/response loo

	// request/response loop
	for {
		// block to read incoming requests
		// since we are using a sessionless proto, each request can
		// potentially go to a different client.  Therefore, the ReadFromXXX
		// operation returns the remote address (saved in raddr)
		// where to send the response.
		_, raddr, err := conn.ReadFromUnix(make([]byte, 48))
		if err != nil {
			fmt.Println("error getting request:", err)
			os.Exit(1)
		}
		// ensure raddr is set
		if raddr == nil {
			fmt.Println("warning: request missing remote addr")
			continue
		}
		// go handle request
		go handleRequest(conn, raddr)
	}
}

// handle incoming requests
func handleRequest(conn *net.UnixConn, addr *net.UnixAddr) {
	// get seconds and fractional secs since 1900
	secs, fracs := getNTPSeconds(time.Now())

	// response packet is filled with the seconds and
	// fractional sec values using Big-Endian
	rsp := make([]byte, 48)
	// write seconds (as uint32) in buffer at [40:43]
	binary.BigEndian.PutUint32(rsp[40:], uint32(secs))
	// write seconds (as uint32) in buffer at [44:47]
	binary.BigEndian.PutUint32(rsp[44:], uint32(fracs))

	// send response to client
	fmt.Printf("writing response %v to %v\n", rsp, addr)
	if _, err := conn.WriteToUnix(rsp, addr); err != nil {
		fmt.Println("err sending data:", err)
		os.Exit(1)
	}
}
