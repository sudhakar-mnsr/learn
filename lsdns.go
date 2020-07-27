package main

import (
   "context"
   "fmt"
   "os"
   "net"
   "flag"
)

var (
   ip string
   host string
   ns bool
   mx bool
   txt bool
   cname bool
)

func init() {
   flag.StringVar(&ip, "ip", "", "IP address for DNS operation")
   flag.StringVar(&host, "host", "", "Host name for DNS operation")
   flag.BoolVar(&ns, "ns", false, "Host name server lookup")
   flag.BoolVar(&mx, "mx", false, "Host domain mail server lookup")
   flag.BoolVar(&txt, "txt", false, "Host domain TXT lookup")
   flag.BoolVar(&cname, "cname", false, "Host CNAME lookup")
}
