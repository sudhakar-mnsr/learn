package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	curr "currency/lib"
)

var (
	currencies = curr.Load("../../../data.csv")
)

// This program implements a simple currency lookup service over 
// TCP or UDS. The request data is unmarshalled to Go type 
// curr.CurrencyRequest using the encoding/json package.

// The request is then used to search the list of currencies. The search 
// result, a []curr.Currency, is marshalled as JSON array of objects and
// sent to the client.

// Focus:
// This version of the program highlights the use of encoding packages
// to serialize data to/from Go data types to another representation 
// such as JSON. The program uses the bufio package to stream data to 
// and from the client. This time however, char '}' is used as demarcation
// instead of '\n'

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
   
   reader := bufio.NewReaderSize(conn, 4)
   
   for {
      // reader will read bytes until '}' is encounter which should indicate
      // the end of JSON object
      buf, err := reader.ReadSlice('}')
      if err != nil {
         if err != io.EOF {
            log.Println("connection read error:", err)
            return
         }
      }
      reader.Reset(conn)
      
      var req curr.CurrencyRequest
      if err := json.Unmarshal(buf, &req); err != nil {
         log.Println("failed to unmarshal request:", err)
         cerr, jerr := json.Marshal(&curr.CurrencyError{Error: err.Error()})
         if jerr != nil {
            log.Println("failed to marshal CurrencyError:", jerr)
            continue
         }
      
         if _, werr := conn.Write(cerr); werr != nil {
            log.Println("failed to write to CurrencyError:", werr)
            return
         }
         continue
      }
      result := curr.Find(currencies, req.Get)
      rsp, err := json.Marshal(&result)
      if err != nil {
         log.Println("failed to marshal data:", err)
         if _, err := fmt.Fprintf(conn, `{"currency_error":"internal error"}`); err != nil {
            log.Printf("failed to write to client: %v", err)
            return
         }
         continue
      }
      
      if _, err := conn.Write(rsp); err != nil {
         log.Println("failed to write response:", err)
         return
      }
   }
}
