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
