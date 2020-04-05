package main

import (
   "fmt"
   "os"
   "net/http"
   "io"
)

func main() {
   client := http.Client{}
   resp, err := client.Get("http://gutenberg.org/cache/epub/16328/pg16328.txt")
   if err != nil {
      fmt.Println(err)
      return
   }
   defer resp.Body.Close()
   
   file, err := os.Create("mnsr.txt")
   if err != nil {
      fmt.Println(err)
      return
   }
   defer file.Close()
   
   io.Copy(file, resp.Body)
   fmt.Println("Text copied to "file.Name())
}
