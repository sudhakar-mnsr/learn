package main

import (
   "fmt"
   "net/http"
   "os"
   "io"
)

func main() {
client := &http.Client{}
req, err := http.NewRequest("GET", "http://tools.ietf.org/rfc/rfc7540.txt", nil)
if err != nil {
   fmt.Println(err)
   return
}

req.Header.Add("Accept", "text/plain")
req.Header.Add("User-Agent", "SampleClient/1.0")

 
   
