package main

import (
"crypto/tls"
"crypto/x509"
"fmt"
"io/ioutil"
"log"
"net/http"
)

func main() {
   // Request /hello over port 8080 via GET method
   // r, err := http.Get("http://localhost:8080/hello")
   // Request /hello over HTTPS port 8443 via the GET method
   // r, err := http.Get("https://localhost:8443/hello")
   // Read the key pair to create certificate
   cert, err := tls.LoadX509KeyPair("/tmp/certs/cert.pem", "/tmp/certs/key.pem")
   if err != nil {
      log.Fatal(err)
   }

   // Create a CA certificate pool and add cert.pem to it
   caCert, err := ioutil.ReadFile("/tmp/certs/cert.pem")
   if err != nil {
      log.Fatal(err)
   }
   caCertPool := x509.NewCertPool()
   caCertPool.AppendCertsFromPEM(caCert)

   // Create a HTTPS cleint and supply the created CA pool
   // Create a HTTPS cleint and supply the created CA pool and certificate
   client := &http.Client{
           Transport: &http.Transport{
           TLSClientConfig: &tls.Config{
              RootCAs: caCertPool,
              Certificates: []tls.Certificate{cert},
           },
        },
   }

   // Request /hello via the created HTTPS client over 8443 via GET
   r, err := client.Get("https://localhost:8443/hello")
   if err != nil {
      log.Fatal(err)
   }
   
   defer r.Body.Close()
   body, err := ioutil.ReadAll(r.Body)
   
   if err != nil {
      log.Fatal(err)
   }
   
   fmt.Printf("%s\n", body)
}
