package main

import (
   "fmt"
   "net"
   "time"
)

var host, port = "127.0.0.1", "4040"
var addr = net.JoinHostPort(host, port)
var deadline = time.Now().Add(time.Millisecond * 700)

const bufLen = 1024
const prompt = "curr"

func main() {
conn, err := net.Dial("tcp", addr)
if err != nil {
   panic(err.Error())
}
defer conn.Close()
fmt.Println("Connected to currency service...")
var cmd, param string

