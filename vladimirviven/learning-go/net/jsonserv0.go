package main

import (
   "encoding/json"
   "fmt"
   "net/http"
   "currency1"
)

// Simple json server for currency service
// As in tcpserver example, returns []currency1.Currency
// This time, however, data is encoded as JSON
// Test with:
// curl-X POST -d '{"get":"Euro"} http://localhost:4040/currency

var currencies = currency1.Load("./data.csv")

func currs(resp http.ResponseWriter, req *http.Request) {
   var currRequest currency1.CurrencyRequest
   dec := json.NewDecoder(req.Body)
   if err := dec.Decode(&currRequest); err != nil {
      resp.WriteHeader(http.StatusBadRequest)
      return
   }
   
   result := currency1.Find(currencies, currRequest.Get)
   enc := json.NewEncoder(resp)
   if err := enc.Encode(&result); err != nil {
      fmt.Println(err)
      resp.WriteHeader(http.StatusInternalServerError)
      return
   }
}

func main() {
   mux := http.NewServeMux()
   mux.HandleFunc("/currency", currs)
   if err := http.ListenAndServe(":4040", mux); err != nil {
      fmt.Println(err)
   }
}
