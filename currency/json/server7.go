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
// search requests JSON objects unmarshalled to curr.CurrencyRequest
// search results curr.Currency marshalled to JSON array and sent back
// Focus:
// This version improve robustness by introducing configuration for read
// and write timeout values. This ensures that client cannot hold a 
// connection hostage by taking long time to send or receive data

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
   
   acceptDelay := time.Millisecond * 10
   acceptCount := 0
    
   for {
      conn, err := ln.Accept()
      if err != nil {
         switch e := err.(type) {
         case net.Error:
            if e.Temporary() {
               if acceptCount > 5 {
                  log.Printf("unable to connect after %d retries: %v", err)
                  return
               }
               acceptDelay *= 2
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
      log.Println("connected to ", conn, RemoteAddr())
      go handleConnection(conn)
   }
}

func handleConnection(conn net.Conn) {
   defer func() {
      if err := conn.Close(); err != nil {
         log.Println("error closing connection:", err)
      }
   }()
   
   // Set initial deadline prior to entering
   // the client request/response loop to 45 seconds.
   // This means the client has 45 seconds to send its initial request or
   // loose the connection.
   
   if err := conn.SetDeadline(time.Now().Add(time.Second * 45)); err != nil {
      log.Println("failed to set deadline:", err)
      return
   }
   
   for {
      dec := json.NewDecoder(conn)
      var req curr.CurrencyRequest
      if err := dec.Decode(&req); err != nil {
         switch err := err.(type) {
         case net.Error:
            // is it a timeout error
            // A deadline policy maybe implemented here using a decreasing
            // grace period that eventually causes an error if reached.
            // Here we just reject the connection if timeout is reached.
            if err.Timeout() {
               fmt.Println("deadline reached, disconnecting...")
            }
            fmt.Println("network error:", err)
            return
         default:
            if err == io.EOF {
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
         switch er := err.(type) {
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
      
      if err := conn.SetDeadline(time.Now().Add(time.Second * 90)); err != nil {
         fmt.Println("failed to set deadline:", err)
         return
      }
   }
} 
