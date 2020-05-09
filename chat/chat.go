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
