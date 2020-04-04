package main

import (
   "fmt"
   "net"
   "strings"
   "time"
   "os"
   curr "currency0"
)

var currencies "./data.csv"

func main() {
   ln, err := net.Listen("tcp", ":4040")
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   defer ln.Close()
   fmt.Println("Global currency service... Listening on port 4040")
   
   for {
      conn, err := ln.Accept()
      if err != nil {
         fmt.Println(err)
         conn.Close()
         continue
      }
      fmt.Println("Connected to ", conn.RemoteAddr())
      // delegate connection to goroutine
      go handleConnection(conn)
   }
}

