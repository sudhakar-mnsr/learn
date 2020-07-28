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


