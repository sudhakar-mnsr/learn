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

   switch network {
   case "udp", "udp4", "udp6", "unixgram":
   default:
      fmt.Println("unsupported network:", network)
      os.Exit(1)
   }

   // create a generic packet connection, PacketConn, with ListenPacket.
   // PacketConn implements common ReadFrom and WriteTo that are 
   // protocol agnostic
   conn, err := net.ListenPacket(network, host)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }
   defer conn.Close()
   fmt.Printf("listening on (%s)%s\n", network, conn.LocalAddr())
