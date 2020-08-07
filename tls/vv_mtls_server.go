package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	curr "currency/lib"
)

var (
	currencies = curr.Load("../../data.csv")
)

func main() {
// setup flags
var addr, network, cert, key, ca string
flag.StringVar(&addr, "e", ":4443", "service endpoint [ip addr or socket path]")
flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
flag.StringVar(&cert, "cert", "../certs/localhost-cert.pem", "public cert")
flag.StringVar(&key, "key", "../certs/localhost-key.pem", "private key")
flag.StringVar(&ca, "ca", "../certs/ca-cert.pem", "root CA certificate")
flag.Parse()

// validate supported network protocols
switch network {
case "tcp", "tcp4", "tcp6", "unix":
default:
   fmt.Println("unsupported network protocol")
   os.Exit(1)
}

// load server cert by providing the private key and generated it.
cer, err := tls.LoadX509KeyPair(cert, key)
if err != nil {
   log.Fatal(err)
}

caCert, err := ioutil.Readfile(ca)
