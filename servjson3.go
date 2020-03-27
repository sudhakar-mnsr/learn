package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	curr "github.com/vladimirvivien/go-networking/currency/lib"
)

var (
	currencies = curr.Load("../data.csv")
)

// This program implements a simple currency lookup service
// over TCP or Unix Data Socket. It loads ISO currency
// information using package curr (see above) and uses a simple
// JSON-encode text-based protocol to exchange data with a client.
//
// Clients send currency search requests as JSON objects
// as {"Get":"<currency name,code,or country"}. The request data is
// then unmarshalled to Go type curr.CurrencyRequest using
// the encoding/json package.
//
// The request is then used to search the list of
// currencies. The search result, a []curr.Currency, is marshalled
// as JSON array of objects and sent to the client.
//
// Focus:
// This version of the program highlights the use of the encoding
// packages to serialize data to/from Go data types to another
// representation such as JSON.  The program uses the bufio package
// to stream data to and from the client as was done with the previous
// version.  This time, however, char '}' is used as demarcation instead
// of '\n'.
//
// Testing:
// Netcat can be used for rudimentary testing.  However, use clientjsonX
// programs functional tests.
//
// Usage: server [options]
// options:
//   -e host endpoint, default ":4040"
//   -n network protocol [tcp,unix], default "tcp"
