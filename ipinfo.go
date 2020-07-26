package main

import (
   "fmt"
   "net"
   "os"
)

func main() {
if len(os.Args) != 2 {
   fmt.Println("Missing ip address")
   os.Exit(1)
}

ip = net.IP(os.Args[1])
if ip == nil {
   fmt.Println("Unable to parse IP address")
   fmt.Println("Address should use IPv4 dot-notation or IPv6 colon-notation")
   os.Exit(1)
} 
}
