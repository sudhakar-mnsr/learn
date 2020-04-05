package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "net/http"
   "currency1"
)

// Currency service client json
// sends request as currency1.CurrencyRequest
// receives []currency1.Currency
func main() {
   var param string
   fmt.Print("Currency> ")
   _, err := fmt.Scanf("%s", &param)
   
   // encode request
   buf := new(bytes.Buffer)
   currRequest := &currency1.CurrencyRequest{Get: param}
   err = json.NewEncoder(buf).Encode(currRequest)
   if err != nil {
      fmt.Println(err)
      return
   }
   
   // send request
   client := &http.Client{}
   req, err := http.NewRequest("POST", "http://127.0.0.1:4040/currency", buf)
   if err != nil {
      fmt.Println(err)
      return
   }
   resp, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer resp.Body.Close()
   var currencies []currency1.Currency
   err = json.NewDecoder(resp.Body).Decode(&currencies)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(currencies)
}
