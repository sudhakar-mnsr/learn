package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	host    string
	network string
)

// This program is a simple Network Time Protocol server that can use
// either UDP or the Unix Domain Socket Datagram protocol.  The program
// uses the ListenPacket to create a PacketConn generic connection.
//
// The server returns the number of seconds since 1900 up to the
// current time. It uses command-line flag -e to specify server
// addr:port and -n to specify network protocol ["udp","unixgram"]

func main() {
   var host, network string
   flag.StringVar(&host,"e",":1123","Server endpoint")
   flag.StringVar(&network,"n","udp","Network protocol")
   
   // validate network protocols
   switch network {
   case "udp", "udp4", "udp6", "unixgram":
   default:
      fmt.Println("unsupported network:", network)
      os.Exit(1)
   }
   
   // create a generic packet connection, PacketConn, with 
   // ListenPacket. PacketConn implements common ReadFrom and
   // WriteTo that are protocol agnostic.
   conn, err := net.ListenPacket(network, host)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }
   defer conn.Close()
   fmt.Println("listening on (%s)%s\n", network, conn.LocalAddr())
   
   for {
      _, raddr, err := conn.ReadFrom(make([]byte, 48))
      if err != nil {
         fmt.Println("error getting request:", err)
         os.Exit(1)
      }
      
      // ensure raddr is set
      if raddr == nil {
         fmt.Println("warning: request missing remote addr")
         continue
      }
      
      go handleRequest(conn, raddr)
   }
}
