package main

import (
   "io"
   "net/http"
   "os"
)

func main() {
   resp, err := http.Get("http://gutenberg.org/cache/epub/16328/pg16328.txt")
   if err != nil {
      panic(err.Error())
   }
   defer resp.Body.Close()
   
   io.Copy(os.Stdout, resp.Body)
}
