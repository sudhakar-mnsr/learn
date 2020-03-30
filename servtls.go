package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	curr "github.com/vladimirvivien/go-networking/tcp/curlib"
)

var (
	currencies = curr.Load("./data.csv")
)

// This program implements a simple currency lookup service
// over TCP or Unix Data Socket. It loads ISO currency
// information using package curlib (see above) and makes
// and serves it using JSON-enocoded data.
//
// Clients send currency search requests as JSON objects such
// as {"Get":"USD"}. The request data is then unmarshalled to Go
// type curr.CurrencyRequest{Get:"USD"} using the encoding/json
// package.
//
// The request is then used to search the list of
// currencies. The search result, a []curr.Currency, is marshalled
// to JSON array of objects and send to the client.
//
// Configure Connection:
// This version of the server highlights the configuration of
// the connection to set read and write deadline for the client.
// If those deadlines are reached, the server will drop the connection.
//
// Usage: server [options]
// options:
//   -e host endpoint, default ":4040"
//   -n network protocol [tcp,unix], default "tcp"
