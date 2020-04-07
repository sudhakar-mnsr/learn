package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

// This program implements a simple echo server over that is able
// to use TCP or Unix Domain Socket (streaming).
// When the server receives a request, it returns its content immediately.
//
// Usage:
// echos2
//   -e <endpoint: ip addr or path>
//   - n <protoco [tcp,unix]>

