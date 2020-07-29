package main

import (
   "encoding/binary"
   "flag"
   "fmt"
   net"
   "os"
   "time"
)   

var (
   host string
   network string
)

// Simple NTP server can use UDP or Unix Domain socket protocol
// ListenPacket to create PacketConn generic connection.
func main() {
flag.StringVar(&host, "e", ":1123", server address
flag.StringVar(&network, "n", "udp", "then network protocol [udp, unixgram]")
flag.Parse

switch network {
   case "udp", "udp4", "udp6", "unixgram":
   default:
      fmt.Println("unsupported network", network
      os.Exit(1)
   }

// Create a generic packet connection, packetConn, with ListenPacket.
conn, err := net.ListenPacket(network, host)
if err != nil {
   fmt.Println("failed to create socket:", err)
   os.Exit(1)
}

defer conn.Close()
fmt.Printf("listening on (%s)%s", network, conn.LocalAddr())

// request response loop
for {
// block to read incoming requests
// since we are using a sessionless protocol, each request can 
// potentially go to a different client. Therefore, RaadFrom operation
// returns remote addres where to send the response

_, raddr, err := conn.ReadFrom(make[]byte, 48)
if err != nil {
   fmt.Println("error getting request:", err
   os.Exit(1)
}

// ensure raddr is set
if raddr == nil {
   fmt.Println("Warning: request missing remote addr")
   continue
}

go handleRequest(conn, raddr)
}
}   

 
