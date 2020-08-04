package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"

	curr "currency/lib"
)

const prompt = "currency"

// It sends JSON-encoded requests and receives JSON array of currency
// information directly over TCP or UDS
// Focus:
// This version of client program highlights the use of IO streaming
// data serilization and client-side error handling

func main() {
   var addr string
   var network string
   flag.StringVar(&addr, "e", "localhost:4040", "service endpoint [ip addr or socket path]")
   flag.StringVar(&network, "n", "tcp", "network protocol [tcp, unix]")
   flag.Parse()
   
   conn, err := net.Dial(network, addr)
   if err != nil {
      fmt.Println("failed to create socket:", err)
      os.Exit(1)
   }
   
   defer conn.Close()
   fmt.Println("connected to currency service:", addr)
   
   var param string
   for {
      fmt.Println("Enter search string or *")
      fmt.Print(promt, ">")
      _, err = fmt.Scanf("%s", &param)
      if err != nil {
         fmt.Println("Usage: <search string or *>")
         continue
      }
      
      req := curr.CurrencyRequest{Get: param}
      // use json encoder to encode value of type curr.CurrencyRequest
      // and stream it to the server via net.Conn
      if err := json.NewEncoder(conn).Encode(&req); err != nil {
         switch err := err.(type) {
         case net.Error:
            fmt.Println("failed to send request:", err)
            os.Exit(1)
         default:
            fmt.Println("failed to encode request:", err)
            continue
         }
      } 
      
      var currencies []curr.Currency
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
      
      for i, c := range currencies {
         fmt.Printf("%2d. %s[%s]\t%s, %s\n", i, c.Code, c.Number, c.Name, c.Country)
      }
   }
}
