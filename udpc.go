package main

import (
   "encoding/binary"
   "fmt"
   "os"
   "flag"
   "net"
   "time"
)

// This program implements NTP client over UDP.
// It uses NTP v3 data packet format ie 48 bits long datagram.

func main() {
var host string
flag.StringVar(&host, "e", "us.pool.ntp.org:123", "NTP host")
flag.Parse()

// req datagram is 48 byte long slice
req := make([]byte, 48)

// req is initialized with 0x1B (request as per protocol) see spec
req[0] = 0x1B

// response 48-byte slice for incoming datagram
rsp := make([]byte, 48)

// Create an address representing remote host
raddr, err := net.ResolveUDPAddr("udp", nil, raddr)
if err != nil {
   fmt.Printf("failed to connect: %v\n", err)
   os.Exit(1)
}

defer func() {
   if err := conn.Close(); err != nil {
      fmt.Println("failed while closing connection:", err)
   }
}()

// setup connection (net.UDPConn) with net.DialUDP
conn, err := net.DialUDP("udp", nil, raddr)
if err != nil {
   fmt.Printf("failed to connect: %v\n", err)
   os.Exit(1)
}

fmt.Printf("time from (udp) %s\n", conn.RemoteAddr())

// Once connection is established, the code pattern
// is the same as in the other impl.

// send the request
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

// ensure we read 48 bytes back (NTP protocol spec)
if read != 48 {
   fmt.Println("did not get all expected bytes from server")
   os.Exit(1)
}


