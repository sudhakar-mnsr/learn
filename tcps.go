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
