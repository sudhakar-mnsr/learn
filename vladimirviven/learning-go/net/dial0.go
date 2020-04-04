package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
   host, port := "www.gutenberg.org", "80"
   addr := net.JoinHostPort(host, port)
   httpRequest := "GET /cache/epub/16328/pg16328.txt HTTP/1.1\n" + "Host: " + host + "\n\n"
   
   // conn, err := net.Dial("udp", addr)
   // The above statement is hanging. When experimenting remember to
   // comment the tcp Dial line below. 
   // Guess: I think this http server at gutenberg could not get
   // port to return as UDP is connection less. Need more clarity
   // on what exactly happens. Then add that detail and remove this 
   // comment.
   // #MNSRremove #MNSRclarity 
   conn, err := net.Dial("tcp", addr)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer conn.Close() 
   
   if _, err = conn.Write([]byte(httpRequest)); err != nil {
      fmt.Println(err)
      return
   }
   
   file, err := os.Create("mnsr.txt")
   if err != nil {
      fmt.Println(err)
      return
   }
   defer file.Close()
   
   io.Copy(file, conn)
   fmt.Println("Text copied to file", file.Name())
}
