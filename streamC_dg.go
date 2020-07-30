package main

import (
   "flag"
   "fmt"
   "os"
   "net"
   "time"
)

// Simple echo client over TCP or unix domain socket.
func main() {
   var addr string
   var network string
   flag.StringVar(&addr, "e", "localhost:4040", "service address endpoint")
   flag.StringVar(&network, "n", "tcp", "network protocol to use")
   flag.Parse()
   text := flag.Arg(0)
   
   switch network {
   case "tcp", "tcp4", "tcp6", "unix":
   default:
      fmt.Println("unsupported network protocol")
      os.Exit(1)
   }
   
   conn, err := net.Dial(network, addr)
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
