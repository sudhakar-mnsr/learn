package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	host    string
	network string
)

// This program is a simple Network Time Protocol server that can use
// either UDP or the Unix Domain Socket Datagram protocol.  The program
// uses the ListenPacket to create a PacketConn generic connection.
//
// The server returns the number of seconds since 1900 up to the
// current time. It uses command-line flag -e to specify server
// addr:port and -n to specify network protocol ["udp","unixgram"]

func main() {
   var host, network string
   flag.StringVar(&host,"e",":1123","Server endpoint")
   flag.StringVar(&network,"n","udp","Network protocol")
   
   // validate network protocols
   switch network {
   case "udp", "udp4", "udp6", "unixgram":
   default:
      fmt.Println("unsupported network:", network)
      os.Exit(1)
   }
   
   // create a generic packet connection, PacketConn, with 
   // ListenPacket. PacketConn implements common ReadFrom and
   // WriteTo that are protocol agnostic.
   conn, err := net.ListenPacket(network, host)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }
   defer conn.Close()
   fmt.Println("listening on (%s)%s\n", network, conn.LocalAddr())
   
   for {
      _, raddr, err := conn.ReadFrom(make([]byte, 48))
      if err != nil {
         fmt.Println("error getting request:", err)
         os.Exit(1)
      }
      
      // ensure raddr is set
      if raddr == nil {
         fmt.Println("warning: request missing remote addr")
         continue
      }
      
      go handleRequest(conn, raddr)
   }
}

// handleRequest handles incoming request and sends current
// time.  If network=udp, the passed address is used.
// If network=unixgram, then the global host address path is
// used for both read and write.
func handleRequest(conn net.PacketConn, addr net.Addr) {
	// get seconds and fractional secs since 1900
	secs, fracs := getNTPSeconds(time.Now())

	// response packet is filled with the seconds and
	// fractional sec values using Big-Endian
	rsp := make([]byte, 48)
	// write seconds (as uint32) in buffer at [40:43]
	binary.BigEndian.PutUint32(rsp[40:], uint32(secs))
	// write seconds (as uint32) in buffer at [44:47]
	binary.BigEndian.PutUint32(rsp[44:], uint32(fracs))

	// send data
	if _, err := conn.WriteTo(rsp, addr); err != nil {
		fmt.Println("err sending data:", err)
		os.Exit(1)
	}

}

// getNTPSecs decompose current time as NTP seconds
func getNTPSeconds(t time.Time) (int64, int64) {
	// convert time to total # of secs since 1970
	// add NTP epoch offets as total #secs between 1900-1970
	secs := t.Unix() + int64(getNTPOffset())
	fracs := t.Nanosecond()
	return secs, int64(fracs)
}
