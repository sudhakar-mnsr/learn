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
