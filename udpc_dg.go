package main

import (
   "encoding/binary"
   "flag"
   "fmt"
   "net"
   "os"
   "time"
)

// NTP client over Unix Domain Socket (datagram)
func main() {
var path string
flag.StringVar(&path, "e", "/tmp/time.sock", "NTP client sock endpoint")
flag.parse()

// req data packet is 48 byte long value
req := make([]byte, 48)

// req is initialized with 0x1B
req[0] = 0x1B

// rsp byte slice used ot receive server response
rsp := make([]byte, 48)

// create a remote address bound to the server socket
raddr, err := net.ResolveUnixAddr("unixgram", path)
if err != nil {
   fmt.Println(err)
   os.Exit(1)
}
