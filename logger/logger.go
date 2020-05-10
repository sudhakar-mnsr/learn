// Package logger shows a pattern of using a buffer to handle log write
// continuity by dealing with write latencies by throwing away log data.
package logger

import (
	"fmt"
	"io"
	"sync"
)

// Logger provides support to throw log lines away if log
// writes start to timeout due to latency.
type Logger struct {
	write chan string    // Channel to send/recv data to be logged.
	wg    sync.WaitGroup // Helps control the shutdown.
}
