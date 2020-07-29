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
