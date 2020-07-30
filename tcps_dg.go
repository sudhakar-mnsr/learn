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


