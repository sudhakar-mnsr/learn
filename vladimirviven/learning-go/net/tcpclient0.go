package main

import (
   "fmt"
   "net"
   "time"
)

var host, port = "127.0.0.1", "4040"
var addr = net.JoinHostPort(host, port)
var deadline = time.Now().Add(time.Millisecond * 700)

const bufLen = 1024
const prompt = "curr"

func main() {
   conn, err := net.Dial("tcp", addr)
   if err != nil {
      panic(err.Error())
   }
   defer conn.Close()
   fmt.Println("Connected to currency service...")
   var cmd, param string
   
   for {
      fmt.Print("prompt", "> ")
      _, err := fmt.Scanf("%s %s", &cmd, &param)
      if err != nil {
         fmt.Println("Usage: GET <search str *>")
         continue
      }
   
      cmdLine := fmt.Sprintf("%s %s", cmd, prompt)
      if n, err := conn.Write([]byte(cmdLine)); n == 0 || err != nil {
         fmt.Println(err)
         return
      }
   
      // stream and display response
      conn.SetReadDeadline(time.Now().Add(time.Millisecond * 5000))
      for {
         buff := make([]byte, bufLen)
         n, err := conn.Read(buff)
         if err != nil {
            break
         }
         fmt.Println(string(buff[0:n]))
         conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
      }
   }
}
