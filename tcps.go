package main

import (
   "flag"
   "fmt"
   "net"
   "os"
)

// Simple TCP echo server
func main() {
   var addr string
   flag.StringVar(&addr, "e", ":4040", "service address endpoint")
   flag.Parse()
   
   laddr, err := net.ResolveTCPAddr("tcp", addr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   
   // anounce service using ListenTCP
   l, err := net.ListenTCP("tcp", laddr)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   defer l.close()
   fmt.Println("listening at (tcp)", laddr.String())
   
   for {
      // use TCPListener to block and wait for TCP
      // connection request using AcceptTCP which creates TCPConn
      conn, err := l.AcceptTCP()
      if err != nil {
         fmt.Println("failed to accept conn:", err)
         conn.Close()
         continue
      }
      fmt.Println("connected to:", conn.RemoteAddr())
      
      go handleConnection(conn)
   }
}

func handleConnection(conn *net.TCPConn) {
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
