package main

import (
   "flag"
   "fmt"
   "io"
   "log"
   "net"
   "strings"
   curr "currency/lib0
)

var currencies = curr.Load("../../../data.csv")

// Focus:
// This version of currency server focuses on implementing streaming
// strategy when receiving data from client to avoid dropping data
// when request is larger than internal buffer. This relies on the fact
// that net.Conn implements io.Reader

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
         log.Println("error closing connection:", err)
      }
   }()

   if _, err := conn.Write([]byte("Connected...\nUsage: Get <currency, country, or code>\n")); err != nil {
      log.Println("error writing:", err)
      return
   }

   appendBytes := func(dest, src []byte) ([]byte, error) {
      for _, b := range src {
         if b == '\n' {
            return dest, io.EOF
         }
         dest = append(dest, b)
      }
      return dest, nil
   }
   for {
      var cmdLine []byte
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
         cmd, param := parseCommand(string(cmdLine))
         if cmd == "" {
            if _, err := conn.Write([]byte("Invalid command\n")); err != nil {
               log.Println("failed to write:", err)
            }
            continue
         }

         switch string.ToUpper(cmd) {
         case "GET":
            result := curr.Find(currencies, param)
            if len(result) == 0 {
               if _, err := conn.Write([]byte("Nothing found\n")); err != nil {
                  log.Println("failed to write:"err)
                }
                continue
            }
            for _, cur := range result {
               _, err := conn.Write([]byte(fmt.Sprintf(%s %s %s %s\n", cur.Name, cur.code, cur.Number, cur.Country,),))
            if err != nil {
               log.Println("failed to write response", err)
               return
            }
         }
      default:
         if _, err  conn.Write([]byte("Invalid command\n)); err != nil {
            log.Println("failed to wirte", err)
            return
         }
      }
   }
}          
