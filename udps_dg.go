package main

import (
   "fmt"
   "encoding/binary"
   "os"
   "net"
   "time"
   "flag"
)

// Simple Network Time Protocol server over Unix Domain Socket
// This implementation uses ListenUnixgram and UnixConn to manage requests.

func main() {
   var path string
   flag.StringVar(&path, "e", "/tmp/time.sock", "NTP server socket endpoint")
   flag.Parse()
   
   // Creates a UnixAddr address
   addr, err := net.ResolveUnixAddr("unixgram", path)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }
   defer conn.Close()
   fmt.Printf("listening on (unixgram) %s\n", conn.LocalAddr())
   
   for {
   // block to read incoming requests
   // since we are using sessionless proto, each request can potentially
   // go to a different client. Therefore the ReadFromXXX operation
   // returns the remote address (saved in raddr) to send response
   
   _, raddr, err := conn.ReadFromUnix(make([]byte, 48))
   if err != nil {
      fmt.Println("error getting request:", err)
      os.Exit(1)
   }
   
   // ensure raddr is set
   if raddr == nil {
      fmt.Println("warning: request missing remote addr")
      continue
   }
   
   //go handleRequest(conn, raddr)
}      

func handleRequest(conn *net.UnixConn, addr *net.UnixAddr) {
   // get seconds and fractional secs since 1900
   secs, fracs := getNTPSeconds(time.Now())
   
   // response packet filled with seconds and
   // fractional sec values using BE
   rsp := make([]byte, 48)
   // write seconds (as uint32) in buffer at [40:43]
   binary.BigEndian.PutUint32(rsp[40:], uint32(secs))
   // write seconds (as uint32) in buffer at [44:47]
   binary.BigEndian.PutUint32(rsp[44:], uint32(fracs))
   
   // send response to client
   fmt.Printf("writing response %v to %v\n", rsp, addr)
   if _, err := conn.WriteToUnix(rsp, addr); err != nil {
      fmt.Println("err sending data:", err)
      os.Exit(1)
   }
}

// getNTPSecs decompose current time as NTP seconds
func getNTPSeconds(t time.Time) (int64, int64) {
   // convert time to total # of secs since 1970
   // add NTP epoch offsets as total #secs between 1900 to 1970
   secs := t.Unix() + int64(getNTPOffset())
   fracs := t.Nanosecond()
   return secs, int64(fracs)
}
