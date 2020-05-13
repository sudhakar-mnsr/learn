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

// New creates a logger value and initializes it for use. The user can
// pass the size of the buffer to use for continuity.
func New(w io.Writer, capacity int) *Logger {

	// Create a value of type logger and init the channel
	// and timer value.
	l := Logger{
		write: make(chan string, capacity), // Buffered channel if size > 0.
	}


	// Add one to the waitgroup to track the write goroutine.
	l.wg.Add(1)

	// Create the write goroutine that performs the actual
	// writes to disk.
