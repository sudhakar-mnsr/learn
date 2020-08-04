package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"learning-go/ch11/curr1"
	"net"
	"os"
	"time"

	curr "currency/lib"
)

const prompt = "currency"

// This program is a client implementation for the currency service
// program. It sends JSON requests and receives JSON-encoded array
// over TCP or UDS
// FOCUS:
// IO streaming
// data serialization
// client-side error handling
// IMPORTANT: configure dialer to setup settings such as timeout and 
// KeepAlive values. Further the code implements a simple connection-retry
// strategy when connecting

func main() {
var addr string
var network string
flag.StringVar(&addr, "e", "localhost:4040", "service endpoint [ip addr or socket path]")
flag.StringVar(&network, "n", "tcp", "network protocol [tcp, unix]")
flag.Parse()

// create a dialer to configure its settings instead of using the default 
// dialer from net.Dial() function
dialer := &net.Dialer{
   Timeout: time.Second * 300,
   KeepAlive: time.Minute * 5,
}

// simple dialing strategy with retry with a simple backoff
// More sophisticated retry strategies
// follow similar pattern but may include features such as exponential
// backoff delay, etc
var (
conn net.Conn
err error
connTries = 0
connMaxRetries = 3
connSleepRetry = time.Second * 1
)

for connTries < connMaxRetries {
fmt.Println("creating connection socket to", addr)
conn, err = dialer.Dial(network, addr)
if err != nil {
   fmt.Println("failed to create socket:", err)
   switch nerr := err.(type) {
   case net.Error:
      if nerr.Temporary() {
         connTries++
         fmt.Println("trying again in:", connSleepRetry)
         time.Sleep(connSleepRetry)
         continue
      }
      fmt.Println("unable to recover")
      os.Exit(1)
   default:
      os.Exit(1)
   }
}
break
}

if conn == nil {
   fmt.Println("failed to create connection successfully")
   os.Exit(1)
}
defer conn.Close()
fmt.Println("connected to currency service: ", addr)

var param string

for {
fmt.Print(prompt, "> ")
_, err = fmt.Scanf("%s", &param)
if err != nil {
   fmt.Println("Usage: <search string or *>")
   continue
}

req := curr.CurrencyRequest{Get: param}

// send request:
// use json encoder to encode value of type curr.CurrencyRequest
// and stream it to the server via net.Conn
if err := json.NewEncoder(conn).Encode(&req); err != nil {
   switch err := err.(type)
   case net.Error:
      fmt.Println("failed to send request:", err)
      os.Exit(1)
   default:
      fmt.Println("failed to encode request:", err)
      continue
   }
}
