package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
        "time"
	curr "currency/lib"
)

var (
	currencies = curr.Load("../../../data.csv")
)

// This program implements simple currency lookup service over TCP
// or UDS. Clients send curency requests as JSON objects such as 
// {"GET":"USD"}. The request data is then unmarshalled to Go type
// curr.CurrencyRequest("GET":"USD"} using encoding/json
// The search result a []curr.Currency, is marshalled to JSON array
// of objects and send to client.
// IO Streaming:
// This version of server highlights the use of IO streaming when
// using net.Conn to stream data to and from clients
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
   fmt.Println("connected to ", conn.RemoteAddr())
   go handleConnection(conn)
   }
}

func handleConnection(conn net.Conn) {
   defer conn.Close()
   // set initial deadline prior to entering the client request/response
   // loop to 90 secs. This means that the client has 90 secs to send its
   // initial request or loose the connection.
   if err := conn.SetDeadline(time.Now().Add(time.Second * 90)); err != nil {
      fmt.Println("failed to set deadline:", err)
      return
   }

   for {
      // The following call uses the JSON encoder support for
      // Go's IO streaming API (io.Reader). It blocks then stream incoming 
      // data from net.Conn implements io.Reader
      dec := json.NewDecoder(conn)
      
      // Next decode the incoming data into Go value curr.CurrencyRequest
      var req curr.CurrencyRequest
      if err := dec.Decode(&req); err != nil {
         // json.Decode() could return decoding err, io err, or networking err
         // hence handle error based on type
         switch err := err.(type) {
         case net.Error:
            fmt.Println("network error:", err)
            return
         default:
            if err == io.EOF {
               fmt.Println("closing connection:", err)
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
      if err := conn.SetDeadline(time.Now().Add(time.Second * 90)); err != nil {
         fmt.Println("failed to set deadline:", err)
         return
      }
   }
}
