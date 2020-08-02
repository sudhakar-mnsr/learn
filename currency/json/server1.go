package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	curr "currency/lib"
)

var (
	currencies = curr.Load("../../../data.csv")
)

// This program implements a simple currency lookup service over
// TCP or UDS. It servers data using JSON encoded data
// Clients send search requests as JSON objects like {"GET":"USD"}
// The request is unmarshalled to Go type curr.CurrencyRequest{Get: "USD"}
// The request is searched and the result a curr.Currency is marshalled
// to JSON array of object and sent to client.
// IO Streaming:
// This version of server highlights the use of IO streaming when
// using net.Conn to stream data to and from clients.

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
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   defer ln.Close()
   fmt.Println("***** Global Currency Service *****")
   fmt.Printf("Service started: (%s) %s\n", network, addr)
   
   for {
      conn, err := ln.Accept()
      if err != nil {
         fmt.Println(err)
         conn.Close()
         continue
      }
      fmt.Println("Connected to ", conn.RemoteAddr())
      go handleConnection(conn)
   }
}

func handleConnection(conn net.Conn) {
   defer conn.Close()
   for {
      // THe following call uses the JSON encoder support for 
      // Go's IO streaming API (io.Reader). It blocks the stream
      // incoming data from net.Conn implements io.Reader.
      dec := json.NewDecoder(conn)

      var req curr.CurrencyRequest
      if err := dec.Decode(&req); err != nil {
         // json.Decode() could return decoding err, io err or networking err
         // so handle err based on type
         switch err := err.(type) {
         case net.Error:
            fmt.Println("network error:", err)
            return
         default:
            if err == io.EOF {
               fmt.Println("closing connection:", err)
               return
            }
            // encode curr.CurrencyError to send to client
            enc := json.NewEncoder(conn)
            if err := enc.Encode(&curr.CurrencyError{Error: err.Error()}); err != nil {
               fmt.Println("failed error encoding:", err)
               return
            }
            continue
         }
      }
      result := curr.Find(currencies, req.Get)
   
      enc := json.NewEncoder(conn)
      if err := enc.Encode(&result); err != nil {
         switch err := err.(type) {
         case net.Error:
            fmt.Println("failed to send response:", err)
            return
         default:
            enc := json.NewEncoder(conn)
            if err := enc.Encode(&curr.CurrencyError{Error: err.Error()}); err != nil {
               fmt.Println("failed to send error:", err)
               return
            }
            continue
         }
      }
   }
} 
