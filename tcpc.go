package main

import (
   "flag"
   "fmt"
   "net"
   "os"
)

// Simple echo client over tcp

func main() {
   var add string
   flag.StringVar(&addr, "e", "localhost:4040", "service address endpoint")
   flag.Parse()
   text := flag.Arg(0)
   
   // use ResolveTCPAddr to create address to connect to
   raddr, err := net.ResolveTCPAddr("tcp", addr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   
   conn, err := net.DialTCP("tcp", nil, raddr)
   if err != nil {
      fmt.Println("failed to connect to server:", err)
      os.Exit(1)
   }
   defer conn.Close()
   
   _, err = conn.Write([]byte(text))
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   
   buf := make([]byte, 1024)
   n, err := conn.Read(buf)
   if err != nil {
      fmt.Println("failed reading response:", err)
      os.Exit(1)
   }
   fmt.Println(string(buf[:n]))
}

