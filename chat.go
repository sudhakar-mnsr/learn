// Package chat implements a basic chat room.
package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

// temporary is declared to test for the existence of the method coming
// from the net package.
type temporary interface {
	Temporary() bool
}
