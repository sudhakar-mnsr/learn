package main

import (
   "crypto/tls"
   "crypto/x509"
   "io/ioutil"
   "io"
   "log"
   "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
   io.WriteString(w, "Hello, world!\n")
}

func main() {
   http.HandleFunc("/hello", helloHandler)
   // Listen to port 8080 and wait 
   // log.Fatal(http.ListenAndServe(":8080", nil))
   // Listen to HTTPS connections on port 8443 and wait
   // log.Fatal(http.ListenAndServeTLS(":8443", "/tmp/certs/cert.pem", "/tmp/certs/key.pem", nil)) 
   // Create a CA certificate pool and add cert.pem to it
   caCert, err := ioutil.ReadFile("/tmp/certs/cert.pem")
   if err != nil {
      log.Fatal(err)
   }
   caCertPool := x509.NewCertPool()
   caCertPool.AppendCertsFromPEM(caCert)
 
   // Create the TLS Config with the CA pool and enable Client certificate
   // validation
   tlsConfig := &tls.Config{
                ClientCAs: caCertPool,
                ClientAuth: tls.RequireAndVerifyClientCert,
   }
   tlsConfig.BuildNameToCertificate()
 
   // Create a server instance to listen on port 8443 with the TLS config
   server := &http.Server{
             Addr: ":8443",
             TLSConfig: tlsConfig,
   }
 
   // Listen to HTTPS connections with the server certificate and wait
   log.Fatal(server.ListenAndServeTLS("/tmp/certs/cert.pem", "/tmp/certs/key.pem"))  
}
