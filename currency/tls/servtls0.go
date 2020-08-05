package main

import (
	"crypto/tls"
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

// Focus:
// tls usage for communication
func main() {
   var addr, network, cert, key string
   flag.StringVar(&addr, "e", ":4443", "service endpoint [ip addr or socket path]")
   flag.StringVar(&network, "n", "tcp", "network protocol [tcp, unix]")
   flag.StringVar(&cert, "cert", "../certs/localhost-cert.pem", "public cert")
   flag.StringVar(&key, "key", "../certs/localhost-key.pem", "private key")
   flag.Parse()
   
   switch network {
   case "tcp", "tcp4", "tcp6", "unix":
   default:
      fmt.Println("unsupported network protocol")
      os.Exit(1)
   }
   // load server cert by providing the private key that generated it
   cer, err := tls.LoadX509KeyPair(cert, key)
   if err != nil {
      log.Fatal(err)
   }
   
   // configure tls with certs and other settings
   tlsConfig := &tls.Config{
      Certificates: []tls.Certificate{cer},
   }
   
   // instead of net.Listen we now use tls.Listen to start
   // a listener on the secure port
   ln, err := tls.Listen(network, addr, tlsconfig)
   if err != nil {
      log.Println(err)
   }
   
   defer ln.Close()
   log.Println("***** Global Currency Service (secure) *****")
   log.Printf("Service started: (%s) %s; server cert %s\n", network, addr, cert)
   
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
                  log.Fatalf("unable to connect after %d retries: %v", acceptCount, err)
               }
               acceptDelay *= 2
               acceptCount++
               time.Sleep(acceptDelay)
               continue
            }
         default:
            log.Println(err)
            if err := conn.Close(); err != nil {
               log.Fatal(err)
            }
            continue
         }
         acceptDelay = time.Millisecond * 10
         acceptCount = 0
      }
      log.Println("securely connected to remote client ", conn.RemoteAddr())
      go handleConnection(conn)
   }
}

func handleConnection(conn net.Conn) {
   defer func() {
      if err := conn.Close(); err != nil {
         log.Println("error closing connection:", err)
      }
   }()
   // set initial deadline prior to enter ing the client request/response
   // loop to 45 seconds. This means that the client has 45 seconds to send
   // its initial request or loose the connection.
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
               log.Println("deadline reached, disconnecting...")
            }
            log.Println("network error:", err)
            return
         default:
            if err == io.EOF {
               log.Println("closing connection:", err)
               return
            }
            enc := json.NewEncoder(conn)
            if encerr := enc.Encode(&curr.CurrencyError{Error: err.Error()}); encerr != nil {
               log.Println("failed error encoding:", encerr)
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
            log.Println("failed to send response:", err)
            return
         default:
            if encerr := enc.Encode(&curr.CurrencyError{Error: err.Error()}); encerr != nil {
               log.Println("failed to send error:", encerr)
               return
            }
            continue
         }
      }
      if err := conn.SetDeadline(time.Now().Add(time.Second * 90)); err != nil {
         log.Println("failed to set deadline:", err)
         return
      }
   }
}
