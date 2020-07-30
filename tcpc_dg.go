package main

import (
"flag"
"net"
"os"
"fmt"
)

// Simple echo server over Unix domain socket

func main() {
   var addr string
   flag.StringVar(&addr, "e", "/tmp/tcpc.sock", "service address endpoint")
   flag.Parse()
   text := flag.Arg(0)
   
   // use ResolveUnixAddr to create remote address to server
   raddr, err := net.ResolveUnixAddr("unix", addr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   
   // use DialUnix to create a connection to remote address.
   // Note: there is no requirement to specify local address
   conn, err := net.DialUnix("unix", nil, raddr)
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
