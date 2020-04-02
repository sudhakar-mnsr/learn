/* FTP Server
 */
package main

import (
        "fmt"
        "net"
        "os"
)

const (
        DIR = "DIR"
        CD  = "CD"
        PWD = "PWD"
)

func main() {

        service := "0.0.0.0:1202"
        tcpAddr, err := net.ResolveTCPAddr("tcp", service)
        checkError(err)

        listener, err := net.ListenTCP("tcp", tcpAddr)
        checkError(err)

        for {
                conn, err := listener.Accept()
                if err != nil {
                        continue
                }
                go handleClient(conn)                                                                              
        }
}

func handleClient(conn net.Conn) {
        defer conn.Close()

        var buf [512]byte
        for {
                n, err := conn.Read(buf[0:])
                if err != nil {
                        conn.Close()
                        return
                }
