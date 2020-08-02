package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	curr "currency/lib"
)

var (
	currencies = curr.Load("../../../data.csv")
)

// This program implements a simple currency lookup service over TCP or
// UDS. Request data is unmarshalled to curr.CurrencyRequest. The response
// a []curr.Currency is marshalled as JSON array of objects and sent

// Focus:
// This version of the program highlights the use of encoding packages
// to serialize data to/from Go data types to another representations
// such as JSON. This version uses the encoding/json package Encoder/Decoder
// types which are accept and io.Writer and io.Reader resp. This means they
// can be used directly with io.Conn value.

func main() {
   // setup flags
   var addr string
   var network string
   flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
   flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
   flag.Parse()
   
   // validate supported network protocols
   switch network {
   case "tcp", "tcp4", "tcp6", "unix":
   default:
   	fmt.Println("unsupported network protocol")
   	os.Exit(1)
   }
   
   ln, err := net.Listen(network, addr)
   log.Println("***** Global Currency Service *****")
   log.Printf("Service started: (%s) %s\n", network, addr)
   
   for {
      conn, err := ln.Accept()
      if err != nil {
         log.Println(err)
         conn.Close()
         continue
      }
      log.Println("Connected to ", conn.RemoteAddr())
      go handleConnection(conn)
   }
}

func handleConnection(conn net.Conn) {
   defer func() {
      if err := conn.Close(); err != nil {
         log.Println("error closing connection:", err)
      }
   }()

   // create json encoder/decoder using the io.Conn as
   // io.Writer and io.Reader for streaming IO
   dec := json.NewDecoder(conn)
   enc := json.NewEncoder(conn)
   for {
      var req curr.CurrencyRequest
      if err := dec.Decode(&req); err != nil {
         log.Println("failed to unmarshal request:", err)
         return
      }
      result := curr.Find(currencies, req.Get)
      if err := enc.Encode(&result); err != nil {
         log.Println("failed to encode data:", err)
         return
      }
   }
} 
