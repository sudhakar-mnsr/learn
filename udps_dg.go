package main

import (
   "fmt"
   "encoding/binary"
   "os"
   "net"
   "time"
   "flag"
)

// Simple Network Time Protocol server over Unix Domain Socket
// This implementation uses ListenUnixgram and UnixConn to manage requests.

func main() {
   var path string
   flag.StringVar(&path, "e", "/tmp/time.sock", "NTP server socket endpoint")
   flag.Parse()
   
   // Creates a UnixAddr address
   addr, err := net.ResolveUnixAddr("unixgram", path)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }
   defer conn.Close()
   fmt.Printf("listening on (unixgram) %s\n", conn.LocalAddr())
   
   for {
   // block to read incoming requests
   // since we are using sessionless proto, each request can potentially
   // go to a different client. Therefore the ReadFromXXX operation
   // returns the remote address (saved in raddr) to send response
   
   _, raddr, err := conn.ReadFromUnix(make([]byte, 48))
   if err != nil {
      fmt.Println("error getting request:", err)
      os.Exit(1)
   }
   
   // ensure raddr is set
   if raddr == nil {
      fmt.Println("warning: request missing remote addr")
      continue
   }
   
   //go handleRequest(conn, raddr)
}      
