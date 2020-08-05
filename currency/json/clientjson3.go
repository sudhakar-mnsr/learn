package main

import (
	"encoding/json"
	"flag"
	"fmt"
	curr1 "currency/lib"
	"net"
	"os"
	"time"

	curr "currency/lib"
)

const prompt = "currency"

// This program is client implementation for currency service
// request is sent in JSON and response is marshalled in JSON array
// This is over TCP or UDS

// Focus:
// This program uses io streaming, data serilization, client-side
// error handling. Configuration of dialer such as timeout, KeepAlive
// Further code also implements connection retry 

func main() {
   var addr string
   var network string
   flag.StringVar(&addr, "e", "localhost:4040", "service endpoint [ip addr or socket path]")
   flag.StringVar(&network, "n", "tcp", "network protocol [tcp,unix]")
   flag.Parse()
   
   // create a dialer to configure its settings instead of using
   // the default dialer from net.Dial() function
   dialer := &net.Dialer{
      Timeout: time.Second * 300,
      KeepAlive: time.Minute * 5,
   }
   
   // simple dialing strategy with retry with a simple backoff.
   // More sophisticated retry strategies follow similar pattern
   // but may include exponential backoff delay
   var (
      conn net.Conn
      err error
      connTries = 0
      connMaxRetries = 3
      connSleepRetry = time.Second * 1
   )
   
   for connTries < connMaxRetries {
      fmt.Println("creating connection socket to ", addr)
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
            // non-recoverable error
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
   fmt.Println("connected to currency service:", addr)
   
   var param string
   
   for {
      fmt.Print(prompt, "> ")
      _, err = fmt.Scanf("%s", &param)
      if err != nil {
         fmt.Println("Usage: <search string or *>")
         continue
      }
      req := curr.CurrencyRequest{Get: param}
      // Send request:
      // use json encoder to encode value of type curr.CurrencyRequest
      // and stream it to the server via net.Conn.
      if err := json.NewEncoder(conn).Encode(&req); err != nil {
         switch err := err.(type) {
         case net.Error:
            fmt.Println("failed to send request:", err)
            os.Exit(1)
         default:
            fmt.Println("failed to encode request", err)
            continue
         }
      }
   
      // Display response
      var currencies []curr1.Currency
      err = json.NewDecoder(conn).Decode(&currencies)
      if err != nil {     
         switch err := err.(type) {
         case net.Error:
            fmt.Println("failed to receive response:", err)
            os.Exit(1)
         default:
            fmt.Println("failed to decode response:", err)
            continue
         }
      }
      fmt.Println(currencies)
   }
}
