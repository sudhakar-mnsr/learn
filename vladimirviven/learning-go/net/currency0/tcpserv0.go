package main

import (
   "fmt"
   "net"
   "strings"
   "time"
   "os"
   curr "currency0"
)

var currencies "./data.csv"

func main() {
   ln, err := net.Listen("tcp", ":4040")
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   defer ln.Close()
   fmt.Println("Global currency service... Listening on port 4040")
   
   for {
      conn, err := ln.Accept()
      if err != nil {
         fmt.Println(err)
         conn.Close()
         continue
      }
      fmt.Println("Connected to ", conn.RemoteAddr())
      // delegate connection to goroutine
      go handleConnection(conn)
   }
}

func handleConn(conn net.Conn) {
   defer conn.Close()
   for {
      cmdLine := make([]byte, 1024 * 4)
      n, err := conn.Read(cmdLine)
      if n == 0 || err != nil {
         return
      }
      cmd, param := parseCommand(string(cmdLine[0:n]))
      if cmd == "" {
         continue
      }
   
      // execute command
      switch strings.ToUpper(cmd)
      case "GET":
         result := curr.Find(currencies, param)
         if len(result) == 0 {
            conn.Write([]byte("Nothing found\n"))
            continue
         }
         // stream result to client
         for _, cur := range result {
            _, err := fmt.Fprintf(conn, "%s %s %s %s\n", cur.Name, cur.Code, cur.Number, cur.Country,)
            if err != nil {
               return
            }
         
            // reset deadline while writing,
            // causes server to close conn if client is gone
            conn.SetWriteDeadline(time.Now().Add(time.second * 5)
         }
         // reset read deadline for next read
         conn.SetWriteDeadline(time.Time{})
      default:
         conn.Write([]byte("Invalid command\n"))
      }
   }
}
      
