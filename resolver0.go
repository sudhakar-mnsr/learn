package main

import (
   "context"
   "flag"
   "fmt"
   "net"
   "os"
)

var (
   host string
)

func main() {
flag.StringVar(&host, "host", "localhost", "host name to resolve")
flag.Parse()

res := net.Resolver{PreferGo: true}
addrs, err := res.LookupHost(context.Background(), host)

