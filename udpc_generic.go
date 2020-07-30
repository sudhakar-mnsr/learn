package main

import (
   "encoding/binary"
   "os"
   "fmt"
   "flag"
   "net"
   "time"
)

// Simple NTP over UDP or Unix Domain Socket datagram
// The program uses Dialer to explicitly configure client dialing process

func main() {
   var host string
   var network string
   flag.StringVar(&host, "e", "us.pool.ntp.org:123", "NTP host")
   flag.StringVar(&network, "n", "udp", "network protocol to use")
   flag.Parse()

   req := make([]byte, 48)
   req[0] = 0x1B

   rsp := make([]byte, 48)

   // Create a Dialer which allows us to specify dialing options.
   // We will need this a bit later to configure the local address
   // when the program is using "unixgram"
   dialer := net.Dialer{}

   // when network is unixgram the local address must be created
   if network == "unixgram" {
      laddr := &net.UnixAddr{Name: fmt.Sprintf("%s-client", host), Net: network}
      dialer.LocalAddr = laddr
   }

   // Setup connection (net.Conn) with Dial()
   conn, err := dialer.Dial(network, host)
   if err != nil {
      fmt.Printf("failed to connect: %v\n", err)
      os.Exit(1)
   }
   defer func() {
      if err := conn.Close(); err != nil {
         fmt.Println("failed while closing connection:", err)
      }
   }()

   // Once connection is established the code pattern is the
   // same as in previous impl
   fmt.Printf("time from (%s) (%s)\n", network, conn.RemoteAddr())

   if _, err = conn.Write(req); err != nil {
      fmt.Printf("failed to send request: %v\n", err)
      os.Exit(1)
   }

   // block to receive server response
   read, err := conn.Read(rsp)
   if err != nil {
      fmt.Printf("failed to receive response: %v\n", err)
      os.Exit(1)
   }

   // ensure we read 48 bytes back (NTP protocol spec)
   if read != 48 {
      fmt.Println("did not get all expected bytes from server")
      os.Exit(1)
   }

   secs := binary.BigEndian.Uint32(rsp[40:])
   frac := binary.BigEndian.Uint32(rsp[44:])

   ntpEpoch := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
   unixEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
   offset := unixEpoch.Sub(ntpEpoch).Seconds()
   now := float64(secs) - offset
   fmt.Printf("%v\n", time.Unix(int64(now), int64(frac)))
}
