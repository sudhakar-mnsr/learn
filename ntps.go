package main

import (
   "encoding/binary"
   "flag"
   "net"
   "fmt"
   "os"
   "time"
)

// This is a simple NTP server over UDP.
// The implementation uses UDPConn and ListenUDP to manage requests.

func main() {
var host string
flag.StringVar(&host, "e", ":1123", "server address")
flag.Parse()

addr, err := net.ResolveUDPAddr("udp", host)
if err != nil {
   fmt.Println(err)
   os.Exit(1)
}

// setup connection UDPConn with ListenUDP
conn, err := net.ListenUDP("udp", addr)
if err != nil {
   fmt.Println("failed to create socket:", err)
   os.Exit(1)
}
defer conn.Close()

fmt.Printf("listening for time requests: (udp) %s\n", conn.LocalAddr())

// read incoming requests.
// since we are using a sessionless proto, each request can
// potentially go to different client. Therefore the ReadFromXXX 
// returns the remote address saved in raddr (to send resp)

_, raddr, err := conn.ReadFromUDP(make([]byte, 48))
if err != nil {
   fmt.Println("error getting request:", err)
   os.Exit(1)
}

// ensure raddr is set
if raddr == nil {
   fmt.Println("request missing remote addr")
   os.Exit(1)
}

// get seconds and fractional secs since 1900
secs, fracs := getNTPSeconds(time.Now())

// response packet is filled with the seconds and
// fractional sec values using Big-Endian
rsp := make([]byte, 48)
// write seconds (as unit32) in buffer at 40:43
binary.BigEndian.PutUint32(rsp[40:], uint32(secs))
// write seconds (as uint32) in buffer at 44:47
binary.BigEndian.PutUint32(rsp[44:], uint32(fracs))

if _, err := conn.WriteToUDP(rsp, raddr); err != nil {
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

// getNTPOffset returns the 70yrs between Unix epoch
// and NTP epoch (1970-1900) in seconds
func getNTPOffset() float64 {
	ntpEpoch := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	unixEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	offset := unixEpoch.Sub(ntpEpoch).Seconds()
	return offset
}

