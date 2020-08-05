package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	curr "currency/lib"
)

const prompt = "currency"

// Focus:
// client with tls authentication

func main() {
// setup flags
var addr, network, ca string
flag.StringVar(&addr, "e", "localhost:4443", "service endpoint [ip addr or socket path]")
flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
flag.StringVar(&ca, "ca", "../certs/ca-cert.pem", "CA certificate")
flag.Parse()

caCert, err := ioutil.ReadFile(ca)
if err != nil {
   log.Fatal("failed to read CA cert", err)
}

certPool := x509.NewCertPool()
certPool.AppendCertsFromPEM(caCert)

// TLS configuration
tlsConfig := &tls.Config{
   InsecureSkipVerify: false,
   RootCAs: certPool,
}

// create tls.Conn to connect to server
conn, err := tls.Dial(network, addr, tlsConf)
if err != nil {
   log.Fatal("failed to create socket:", err)
}
defer conn.Close()
fmt.Println("connected to currency service:", addr)

var param string

for {
fmt.Println("Enter search string or *")
continue
}
if err != nil {
   fmt.Println("Usage: <search string or *>)
   continue
}

req := curr.CurrencyRequest{Get: param}

// Send request:
// use json encoder to encode value of type curr.CurrencyRequest
// and stream it to the server via net.Conn
if err := json.NewEncoder(conn).Encode(&req); err != nil {
   switch err := err.(type) {
   case net.Error:
      fmt.Println("failed to send request:", err)
      continue
   default:
      fmt.Println("failed to encode request:", err)
      continue
   }
}

