package main

import (
   "encoding/binary"
   "flag"
   "net"
   "fmt"
   "os"
   "time"
)

// This is a simple NTP server over UDP.
// The implementation uses UDPConn and ListenUDP to manage requests.

func main() {
var host string
flag.StringVar(&host, "e", ":1123", "server address")
flag.Parse()

addr, err := net.ResolveUDPAddr("udp", host)
if err != nil {
   fmt.Println(err)
   os.Exit(1)
}

