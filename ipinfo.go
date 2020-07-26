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

fmt.Println()
fmt.Println("IP:                     %s\n", ip)
fmt.Println("Default Mask:           %s\n", net.IP(ip.DefaultMask()))
fmt.Println("Loopback:               %t\n", ip.IsLoopback())
fmt.Println("Unicast:")
fmt.Println("   Global:              %t\n", ip.IsGlobalUnicast())
fmt.Println("   Link:                %t\n", ip.IsLinkLocalUnicast())
fmt.Println("Multicast:")
fmt.Println("   Global:              %t\n", ip.IsMulticast())
fmt.Println("   Interface:           %t\n", ip.IsInterfaceLocalMulticast())
fmt.Println("   Link:                %t\n", ip.IsLinkLocalMulticast())
fmt.Println()
}
