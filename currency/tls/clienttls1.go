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
flag.StringVar(&cert, "cert", "../certs/client-cert.pem", "public cert")
flag.StringVar(&key, "key", "../certs/client-key.pem", "private key")
flag.StringVar(&ca, "ca", "../certs/ca-cert.pem", "root CA certificate")
flag.Parse()

cer, err := tls.LoadX509KeyPair(cert, key)
if err != nil {
   log.Fatal(err)
}

caCert, err := ioutil.ReadFile(ca)
if err != nil {
   log.Fatal("failed to read CA cert", err)
}

certPool := x509.NewCertPool()
certPool.AppendCertsFromPEM(caCert)

tlsConf := &tls.Config{
   RootCAs: certPool,
   Certificates: []tls.Certificate{cer},
}
