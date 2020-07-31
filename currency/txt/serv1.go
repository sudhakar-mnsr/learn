package main

import (
"flag"
"fmt"
"log"
"net"
"strings"
curr "github.com/sudhakar-mnsr/currency/lib0"
)

var currencies = curr.Load("../data.csv")

// Building simple text based lookup service over TCP or unix domain socket.
// Uses text based protocol to interact with client and send data
// Protocol: GET <currency, country or code>
// Focus:
// This version of server uses TCP or Unix Domain sockets.
// This is text based application protocol.
// No streaming strategy employed for read/write operations.
// Buffers are read in one shot (chances for missing data during read).
// Testing:
// Netcat or telnet can be used to rest this server by connecting and 
// sending command described above

func main() {
   var addr string
   var network string
   flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
   flag.StringVar(&network, "n", "tcp", "network protocol[tcp, unix]")
   flag.Parse()
   
   switch network {
   case "tcp", "tcp4", "tcp6", "unix":
   default:
      fmt.Println(log.Fatalln("unsupported network protocol:", network)
   }
   
   // create a listener for provided network and host address
   ln, err := net.Listen(network, addr)
   if err != nil {
      log.Fatal("failed to create listener:", err)
   }
   defer ln.Close()
   
   for {
      conn, err := ln.Accept()
      if err != nil {
         fmt.Println(err)
         if err := conn.Close(); err != nil {
            log.Println("failed to close listener:", err)
         }
         continue
      }
      log.Println("Connected to", conn.RemoteAddr())
      
      go handleConnection(conn)
   }
}
