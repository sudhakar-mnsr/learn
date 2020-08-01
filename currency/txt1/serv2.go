package main

import (
	"flag"
	"fmt"
        "io"
	"log"
	"net"
	"strings"

	curr "currency/lib0"
)

var currencies = curr.Load("../../../data.csv")

// This is simple currency lookup service over TCP or Unix Data Socket
// Text based protocol designed to work on top of TCP or UDS
// Focus:
// There is streaming strategy for read/write operations to avoid dropping
// data when the request is larger than internal buffer. This relies on
// fact that net.Conn implements io.Reader which allows stream data

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
   // appendBytes is a function that simulates end-of-file marker error.
   // Since we will use streaming IO on top of a streaming protocol,
   // there may never be an actual EOF marker. so this function simulates
   // io.EOF using char '\n'
   appendBytes := func(dest, src []byte) ([]byte, error) {
      for _, b := range src {
         if b == '\n' {
            return dest, io.EOF
         }
         dest = append(dest, b)
      }
   }

   for {
      var cmdLine []byte
      // stream data using 4-byte chunks until io.EOF
      // The chunks are kept small to demo streaming using io.Reader
      for {
         chunk := make([]byte, 4)
         n, err := conn.Read(chunk)
         if err != nil {
            if err == io.EOF {
               cmdLine, _ = appendBytes(cmdLine, chunk[:n])
               break
            }
            log.Println("connection read error:", err)
            return
         }
         if cmdLine, err = appendBytes(cmdLine, chunk[:n]); err == io.EOF {
            break
         }
      }
      cmd, param := parseCommand(string(cmdLine))
      if cmd == "" {
         if _, err := conn.Write([]byte("Invalid command\n")); err != nil {
            log.Println("failed to write:", err)
            return
         }
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

func parseCommand(cmdLine string) (cmd, param string) {
   parts := strings.Split(cmdLine, " ")
   if len(parts) != 2 {
      return "", ""
   }
   cmd = strings.TrimSpace(parts[0])
   param = strings.TrimSpace(parts[1])
   return
}
