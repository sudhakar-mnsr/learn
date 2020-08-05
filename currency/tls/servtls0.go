package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	curr "currency/lib"
)

var (
	currencies = curr.Load("../../../data.csv")
)

// Focus:
// tls usage for communication
func main() {
var addr, network, cert, key, string
flag.StringVar(&addr, "e", ":4443", "service endpoint [ip addr or socket path]")
flag.StringVar(&network, "n", "tcp", "network protocol [tcp, unix]")
flag.StringVar(&cert, "cert", "../certs/localhost-cert.pem", "public cert")
flag.StringVar(&key, "key", "../certs/localhost-key.pem", "private key")
flag.Parse()

switch network {
case "tcp", "tcp4", "tcp6", "unix":
default:
   fmt.Println("unsupported network protocol")
   os.Exit(1)
}
// load server cert by providing the private key that generated it
cer, err := tls.LoadX509KeyPair(cert, key)
if err != nil {
   log.Fatal(err)
}

// configure tls with certs and other settings
tlsConfig := &tls.Config{
   Certificates: []tls.Certificate{cer},
}

