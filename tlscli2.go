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

	curr "github.com/vladimirvivien/go-networking/currency/lib"
)

const prompt = "currency"
