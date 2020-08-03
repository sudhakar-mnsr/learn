package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	curr "currency/lib"
)

var (
	currencies = curr.Load("../../../data.csv")
)

// This program implements simple currency lookup service over TCP or UDS
// The request data is unmarshalled to curr.CurrencyRequest
// The search result is marshalled as JSON array and sent to client.
// Focus:
// This code is more roboust in parsing network error and implement retry
// On temporary failures the server attempt to retry connecting.

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
      log.Println(err)
      os.Exit(1)
   }
   defer ln.Close()
   log.Println("***** Global Currency Service *****")
   log.Printf("Service started: (%s) %s\n", network, addr)
   
   // delay to sleep when accept fails with a temporary error
   acceptDelay := time.Millisecond * 10
   acceptCount := 0
   
   for {
      conn, err := ln.Accept()
      if err != nil {
         switch e := err.(type) {
         case net.Error:
            if e.Temporary() {
               if acceptCount > 5 {
                  log.Printf("unable to connect after %d retries: %v", acceptCount, err)
                  return
               }
               acceptDelay *=2
               acceptCount++
               time.Sleep(acceptDelay)
               continue
            }
         default:
            log.Println(err)
            conn.Close()
            continue
         }
         acceptDelay = time.Millisecond * 10
         acceptCount = 0
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
   
   for {
      dec := json.NewDecoder(conn)
      var req curr.CurrencyRequest
      if err := dec.Decode(&req); err != nil {
         switch err := err.(type) {
         case net.Error:
            fmt.Println("network error:", err)
            return
         default:
            if err != io.EOF {
               fmt.Println("closing connection:", err)
               return
            }
            enc := json.NewEncoder(conn)
            if encerr := enc.Encode(&curr.CurrencyError{Error: err.Error()}); encerr != nil {
               fmt.Println("failed error encoding:", encerr)
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
            if encerr := enc.Encode(&curr.CurrencyError{Error: err.Error()}); encerr != nil {
               fmt.Println("failed to send error:", encerr)
               return
            }
            continue
         }
      }
   }
}
