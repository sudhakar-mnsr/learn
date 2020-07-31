package main

import (
   "bufio"
   "flag"
   "fmt"
   "io"
   "log"
   "net"
   "strings"
   curr "currency/lib0
)

var currencies = curr.Load("../../../data.csv)

// Focus:
// This version uses the bufio package to use buffered readers
// to stream from net.Conn

func main() {
   var addr string
   var network string
   flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
   flag.StringVar(&network, "n", "tcp", "network protocol [tcp, unix]")
   flag.Parse()
   
   switch network {
   case "tcp", "tcp4", "tcp6", "unix":
   default:
      log.Fatalln("unsupported network protocol:", network)
   }
   
   ln, err := net.Listen(network, addr)
   if err != nil {
      log.Fatal("failed to create listener:", err)
   }
   defer ln.Close()
   log.Println("***** Global Currency Service *****")
   log.Printf("Service started: (%s) %s\n", network, addr)
   
   for {
      conn, err := ln.Accept
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
