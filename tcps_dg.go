package main

import (
"flag"
"fmt"
"net"
"os"
)

// Simple echo server over Unix Domain socket (streaming)

func main() {
   var addr string
   flag.StringVar(&addr, "e", "/tmp/tcps.sock", "service endpoint address")
   flag.Parse()
   
   laddr, err := net.ResolveUnixAddr("unix", addr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   
   // announce service using LIstenUnix which creates UnixListener
   l, err := net.ListenUnix("unix", laddr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   defer l.Close()
   fmt.Println("listening at (unix)", laddr.String())
   
   for {
      // use UnixListener to block and wait for UDS connection request
      // using AcceptUnix which then creates UnixConn
      conn, err := l.AcceptUnix()
      if err != nil {
         fmt.Println("failed to accept conn:", err)
         conn.Close()
         continue
      }
      fmt.Println("connected to: ", conn.RemoteAddr())
   
      go handleConnection(conn)
   }
}

func handleConnection(conn *net.UnixConn) {
   defer conn.Close()
   
   buf := make([]byte, 1024)
   
   n, err := conn.Read(buf)
   if err != nil {
      fmt.Println(err)
      return
   }
   
   w, err := conn.Write(buf[:n])
   if err != nil {
      fmt.Println("failed to write to client:", err)
      return
   }
   
   if w != n {
      fmt.Println("warning: not all data sent to client")
      return
   }
}
