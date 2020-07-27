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

// This program looksup IP address associated with the hostname
// This program uses net.Resolver directly to specify resolver
// net.Resolver{PureGo:true}

func main() {
   flag.StringVar(&host, "host", "localhost", "host name to resolve")
   flag.Parse()
   
   res := net.Resolver{PreferGo: true}
   addrs, err := res.LookupHost(context.Background(), host)
   
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   
   fmt.Println(addrs)
}
