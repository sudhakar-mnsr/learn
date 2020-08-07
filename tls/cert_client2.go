package main

import (
"fmt"
"io/ioutil"
"log"
"net/http"
)

func main() {
   // Request /hello over port 8080 via GET method
   // r, err := http.Get("http://localhost:8080/hello")
   // Request /hello over HTTPS port 8443 via the GET method
   r, err := http.Get("https://localhost:8443/hello")
   // Create a CA certificate pool and add cert.pem to it
   caCert, err := ioutil.ReadFile("cert.pem"
   if err != nil {
      log.Fatal(err)
   }
   caCertPool := x509.NewCertPool()
   caCertPool.AppendCertsFromPEM(caCert)
   
   defer r.Body.Close()
   body, err := ioutil.ReadAll(r.Body)
   
   if err != nil {
      log.Fatal(err)
   }
   
   fmt.Printf("%s\n", body)
}
