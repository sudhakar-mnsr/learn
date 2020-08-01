package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	curr "currency/lib0"
)

var currencies = curr.Load("../../../data.csv")

// This is simple currency lookup service over TCP or Unix Data Socket
// Text based protocol designed to work on top of TCP or UDS
// Focus:
// There is no streaming strategy for read/write operations
// Buffers are read in one shot creating opportunities for missing data

func main() {
   var addr string
   var network string
   flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
   flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
   flag.Parse()
   
   // validate supported network protocols
   switch network {
   case "tcp", "tcp4", "tcp6", "unix":
   default:
   	log.Fatalln("unsupported network protocol:", network)
   }
   
   // create listener for provided network and host address
   ln, err := net.Listen(network, addr)
   if err != nil {
      log.Fatal("failed to create listener:", err)
   }
   defer ln.Close()
   log.Println("***** Global Currency Service *****")
   log.Printf("Service started: (%s) %s\n", network, addr)
   
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

func handleConnection(conn net.Conn) {
   defer func() {
      if err := conn.Close(); err != nil {
         log.Println("error causing connection:", err)
      }
   }()
   if _, err := conn.Write([]byte("Connected...\nUsage: GET <currency, country, or code>\n")); err != nil {
      log.Println("error writing:", err)
      return
   }   
   
   for {
      cmdLine := make([]byte, (1024 * 4))
      n, err := conn.Read(cmdLine)
      if n == 0 || err != nil {
         log.Println("connection read error:", err)
         return
      }
      cmd, param := parseCommand(string(cmdLine[0:n]))
      if cmd == "" {
         if _, err := conn.Write([]byte("Invalid command\n")); err != nil {
            log.Println("failed to write:", err)
            return
         }
         continue
      }
      
      switch strings.ToUpper(cmd) {
      case "GET":
         result := curr.Find(currencies, param)
         if len(result) == 0 {
            if _, err := conn.Write([]byte("Nothing found\n")); err != nil {
               log.Println("failed to write:", err)
            }
            continue
         }
         for _, cur := range result {
            _, err := conn.Write([]byte(fmt.Sprintf("%s %s %s %s\n", cur.Name, cur.Code, cur.Number, cur.Country,),))
            if err != nil {
               log.Println("failed to write response:", err)
               return
            }
         }
      default:
         if _, err := conn.Write([]byte("Invalid command\n")); err != nil {
            log.Println("failed to write:", err)
            return
         }
      }
   }
}
