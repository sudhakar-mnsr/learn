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

func main() {
   // setup flags
   var addr, network, cert, key, ca string
   flag.StringVar(&addr, "e", "localhost:4443", "service endpoint [ip addr or socket path]")
   flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
   flag.StringVar(&cert, "cert", "/tmp/certs/cert.pem", "public cert")
   flag.StringVar(&key, "key", "/tmp/certs/key.pem", "private key")
   flag.StringVar(&ca, "ca", "/tmp/certs/cert.pem", "root CA certificate")
   flag.Parse()
   
   cer, err := tls.LoadX509KeyPair(cert, key)
   if err != nil {
      log.Fatal(err)
   }
   
   // Load our CA certificate
   caCert, err := ioutil.ReadFile(ca)
   if err != nil {
      log.Fatal("failed to read CA cert", err)
   }
   
   certPool := x509,NewCertPool()
   certPool.AppendCertsFromPEM(caCert)
   
   // TLS configuration
   tlsConf := &tls.Config{
           RootCAs: certPool,
           Certificates: []tls.Certificate{cer}.
   }
   
   // create a tls.Conn to connect to server
   conn, err := tls.Dial(network, addr, tlsConf)
   if err != nil {
      log.Fatal("failed to create socket:", err)
   }
   defer conn.Close()
   fmt.Println("connected to currency service:", addr)
   
   var param string
   
   for {
      fmt.Println("Enter search string or *")
      fmt.Print(prompt, ">")
      _, err = fmt.Scanf("%s", &param)
      if err != nil {
         fmt.Println("Usage: <search sting or *")
         continue
      }
      req := curr.CurrencyRequest{Get: param}
      
      // Send request
      // use json encoder to encode value of type curr.CurrencyRequest and stream
      // it to the server via net.Conn 
      
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
      
      // Receive response
      var currencies []curr.Currency
      err = json.NewDecoder(conn).Decode(&currencies)
      if err != nil {
         switch err := err.(type) {
         case net.Error:
            fmt.Println("failed to receive response:", err)
            continue
         default:
            fmt.Println("failed to decode response:, err)
            continue
         }
      }
      for i, c := range currencies {
         fmt.Printf("%2d, %s[%s]\t%s, %s\n", i, c.Code, c.Number, c.Name, c.Country)
      }
   }
}    
