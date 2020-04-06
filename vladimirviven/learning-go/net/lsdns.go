package main

import (
   "context"
   "flag"
   "fmt"
   "net"
   "os"
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
   flag.StringVar(&host, "host", "", "Host address for DNS operation")
   flag.BoolVar(&ns, "ns", false,"Host name server lookup")
   flag.BoolVar(&mx, "mx", false,"Host domain mail server lookup")
   flag.BoolVar(&txt, "txt", false,"Host domain TXT lookup")
   flag.BoolVar(&cname, "cname", false,"Host CNAME lookup")
}

type lsdns struct {
   resolver *net.Resolver
}

func newLsdns() *lsdns {
   return &lsdns{net.DefaultResolver}
}

func (ls *lsdns) reverseLkp(ip string) error {
   names, err := ls.resolver.LookupAddr(context.Background(), ip)
   if err != nil {
      return err
   }
   fmt.Println("Reverse lookup")
   fmt.Println("--------------")
   for _, name := range names {
      fmt.Println(name)
   }
   return nil
}

