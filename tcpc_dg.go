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
flag.StringVar(&host, "e", "/tmp/tcpc.sock", "service address endpoint")
flag.Parse()
text := flag.Arg(0)

// use ResolveUnixAddr to create remote address to server
raddr, err := net.ResolveUnixAddr("unix", addr)
if err != nil {
   fmt.Println(err)
   os.Exit(1)
}

