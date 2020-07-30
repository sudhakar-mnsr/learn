package main

import (
   "encoding/binary"
   "fmt"
   "os"
   "flag"
   "net"
   "time"
)

var (
   host string
   network string
)

// Simple NTP server that can use either UDP or Unix Domain Socket
// This program uses ListenPacket to create PacketConn generic connection

func main() {
   flag.StringVar(&host, "e", ":1123", "server address")
   flag.StringVar(&network, "n", "udp", "udp or unixgram")
   flag.Parse()


