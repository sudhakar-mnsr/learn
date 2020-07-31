package main

import (
"flag"
"fmt"
"log"
"net"
"strings"
curr "github.com/sudhakar-mnsr/currency/lib0"
)

var currencies = curr.Load("../data.csv")

// Building simple text based lookup service over TCP or unix domain socket.
// Uses text based protocol to interact with client and send data
// Protocol: GET <currency, country or code>
// Focus:
// This version of server uses TCP or Unix Domain sockets.
// This is text based application protocol.
// No streaming strategy employed for read/write operations.
// Buffers are read in one shot (chances for missing data during read).
// Testing:
// Netcat or telnet can be used to rest this server by connecting and 
// sending command described above

func main() {
var addr string
var network string
flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or socket path]")
flag.StringVar(&network, "n", "tcp", "network protocol[tcp, unix]")
flag.Parse()


