package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

// This program implements a simple echo client over unix domain socket.
// It sends a text content to the server and displays
// the response on the screen.
//
// Usage: echoc -e <host-addr-endpoint> <text content>
