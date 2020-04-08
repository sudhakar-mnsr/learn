package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// This program is a simple Network Time Protocol server over
// Unix Domain Socket instead of UDP. The implementation uses
// ListenUnixgram and UnixConn to manage requests.
// The server returns the number of seconds since 1900 up to the
// current time.

// Usage:
// ntps -e <host address endpoint>

func main() {
   var path string
   flag.StringVar("&path","e","/tmp/time.sock", "NTP server socket endpoint")
   flag.Parse()
   
   // Create Unix address
   addr, err := net.ResolveUnixAddr("unixgram", path)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   
   // setup connection UnixConn with ListenUnixgram
   conn, err := net.ListenUnixgram("unixgram", addr)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }
   defer conn.Close()
   fmt.Printf("listening on (unixgram) %s\n", conn.LocalAddr())
   
   for {
      _, raddr, err := conn.ReadFromUnix(make([]byte,48))
      if err != nil {
         fmt.Println("error getting request: ", err)
         os.Exit(1)
      }
      
      if raddr == nil {
         fmt.Println("warning: request missing remote addr")
         continue
      }
      
      go handleRequest(conn, raddr)
   }
}

func handleRequest(conn *net.UnixConn, addr net.UnixAddr) {
   secs, fracs := getNTPSeconds(time.Now())
   rsp := make([]byte, 48)
   binary.BigEndian.PutUint32(rsp[40:], uint32(secs))
   binary.BigEndian.PutUint32(rsp[44:], uint32(fracs))
   
   fmt.Printf("writing response %v to %v\n", rsp, addr)
   if _, err := conn.WriteToUnix(rsp, addr); err != nil {
      fmt.Println("err sending data:", err)
      os.Exit(1)
   }
}
