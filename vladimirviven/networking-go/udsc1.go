package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// This program implements an NTP client that is capable of
// using either UDP or Unix Domain Socket datagram.  To do this,
// the program uses the Dialer to explicitly configure the client
// dialing process.
//
// The program uses -host to specify the remote address
// (or socket path) and -n for the network protocl ("udp" or "datagram").

func main() {
var host, network string
flag.StringVar(&host, "e",":1123","Service Endpoint")
flag.StringVar(&network, "n","udp","Network protocol")
flag.Parse()

// Create a Dialer which allows us to specify dialing options
// we will need this a bit later to configure the local address
// when the program is using "unixgram"
dialer := net.Dialer{}

if network == "unixgram" {
   laddr := &net.UnixAddr{Name: fmt.Sprintf("%s-client", host),
                          Net: network}
   dialer.LocalAddr := laddr
} 

conn, err := dialer.Dial(network, host)
if err != nil {
   fmt.Printf("failed to connect: %v\n", err)
   os.Exit(1)
}
defer func() {
   if err := conn.Close(); err != nil {
      fmt.Println("failed to close connection", err)
   }
}()

// Once the connection is established, the code pattern
// is the same as in the previous impl.

fmt.Printf("time from (%s) (%s)\n", network, conn.RemoteAddr())

// Send time request
if _, err = conn.Write(req); err != nil {
   fmt.Printf("failed to send request: %v\n", err)
   os.Exit(1)
}

// block to recieve server response
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

	// NTP data comes in as big-endian (LSB [0...47] MSB)
	// with a 64-bit value containing the server time in seconds
	// where the first 32-bits are seconds and last 32-bit are fractional.
	// The following extracts the seconds from [0...[40:43]...47]
	// it is the number of secs since 1900 (NTP epoch)
	secs := binary.BigEndian.Uint32(rsp[40:])
	frac := binary.BigEndian.Uint32(rsp[44:])
