package main

import (
   "encoding/binary"
   "fmt"
   "os"
   "flag"
   "net"
   "time"
)

var (
   host string
   network string
)

// Simple NTP server that can use either UDP or Unix Domain Socket
// This program uses ListenPacket to create PacketConn generic connection

func main() {
   flag.StringVar(&host, "e", ":1123", "server address")
   flag.StringVar(&network, "n", "udp", "udp or unixgram")
   flag.Parse()

   switch network {
   case "udp", "udp4", "udp6", "unixgram":
   default:
      fmt.Println("unsupported network:", network)
      os.Exit(1)
   }

   // create a generic packet connection, PacketConn, with ListenPacket.
   // PacketConn implements common ReadFrom and WriteTo that are 
   // protocol agnostic
   conn, err := net.ListenPacket(network, host)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }
   defer conn.Close()
   fmt.Printf("listening on (%s)%s\n", network, conn.LocalAddr())

   for {
      // use generic ReadFrom instead of ReadFromXXX
      // ReadFrom returns remote address (which addr udp or unixgram)
      // MNSR provide clarity on the above (how it knows)
      _, raddr, err := conn.ReadFrom(make([]byte, 48))
      if err != nil {
         fmt.Println("error getting request:", err)
         os.Exit(1)
      }

      if raddr == nil {
         fmt.Println("warning: request missing remote addr")
         continue
      }

      go handleRequest(conn, raddr)
   }
}

// network=udp passed address is used.
// network=unixgram the global host address path is used for read/write
func handleRequest(conn net.PacketConn, addr net.Addr) {
   secs, fracs := getNTPSeconds(time.Now())
   rsp := make([]byte, 48)
   
   // write seconds as uint32 in buffer at [40:43]
   binary.BigEndian.PutUint32(rsp[40:], uint32(secs))
   // write seconds as uint32 in buffer at [44:47]
   binary.BigEndian.PutUint32(rsp[44:], uint32(fracs))
   
   if _, err := conn.WriteTo(rsp, addr); err != nil {
      fmt.Println("err sending data:", err)
      os.Exit(1)
   }
}

// getNTPSecs decompose current time as NTP seconds
func getNTPSeconds(t time.Time) (int64, int64) {
   // convert time to total # of secs since 1970
   // add NTP epoch offsets as total #secs between 1900-1970
   secs := t.Unix() + int64(getNTPOffset())
   fracs := t.Nanosecond()
   return secs, int64(fracs)
}

// getNTPOffset returns the 70yrs between unix epoch
// and NTP epoch (1970-1900) in seconds
func getNTPOffset() float64 {
   ntpEpoch := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
   unixEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
   offset := unixEpoch.Sub(ntpEpoch).Seconds()
   return offset
}
