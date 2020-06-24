// Sample program to show how the goroutine scheduler
// will time slice goroutines on a single thread.
package main

import (
	"crypto/sha1"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)
