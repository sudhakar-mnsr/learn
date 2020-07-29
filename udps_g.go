package main

import (
   "encoding/binary"
   "flag"
   "fmt"
   net"
   "os"
   "time"
)   

var (
   host string
   network string
)

// Simple NTP server can use UDP or Unix Domain socket protocol
// ListenPacket to create PacketConn generic connection.
func main() {
   flag.StringVar(&host, "e", ":1123", server address
   flag.StringVar(&network, "n", "udp", "then network protocol [udp, unixgram]")
   flag.Parse
   
   switch network {
      case "udp", "udp4", "udp6", "unixgram":
      default:
         fmt.Println("unsupported network", network
         os.Exit(1)
      }
   
   // Create a generic packet connection, packetConn, with ListenPacket.
   conn, err := net.ListenPacket(network, host)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }
   
   defer conn.Close()
   fmt.Printf("listening on (%s)%s", network, conn.LocalAddr())
   
   // request response loop
   for {
   // block to read incoming requests
   // since we are using a sessionless protocol, each request can 
   // potentially go to a different client. Therefore, RaadFrom operation
   // returns remote addres where to send the response
   
   _, raddr, err := conn.ReadFrom(make[]byte, 48)
   if err != nil {
      fmt.Println("error getting request:", err
      os.Exit(1)
   }
   
   // ensure raddr is set
   if raddr == nil {
      fmt.Println("Warning: request missing remote addr")
      continue
   }
   
   go handleRequest(conn, raddr)
   }
}   

// handleRequest handles incoming request and sends current time.
// if network=unixgram the the global host address path is used
// for both read and write
func handleRequest(conn net.PacketConn, addr net.Addr) {
   // get seconds and fractional secs since 1900
   secs, fracs :=getNTPSeconds(time.now())
   
   // response packet is filled with the seconds and the fractional
   // sec values using Big-Endian
   
   rsp := make([]byte, 48)
   // write seconds (as uingt32) in buffer at [40:43] 
   binary.BigEndian.PutUint32(rsp[40:], uint32(secs))
   // Write seconds as uint32 in buffer at [44:47]
   binary.BigEndian.PutUnit32(rsp[44:], uint32(fracs))
   
   // Send data
   if _, err := conn.writeTo(rsp, addr); err != nil {
      fmt.Println("err sending data:", err)
      os.Exit(1)
}

func getNTPSeconds(t time.Time) (int64, int64) {
   // convert time to total # of secs since 1970
   // add NTP offsets as total #secs between 1900 to 1970
   secs := t.Unix() + int64(getNTPOffset())
   fracs := t.Nanosecond()
   return secs, int64(fracs)
}
