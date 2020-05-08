
// Sample program to see what a trace will look like for basic
// channel latencies.
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// data represents a set of bytes to process.
var data []byte

// init creates a data for processing.
func init() {
	f, err := os.Open("data.bytes")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data, err = ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bytes", len(data))
}

// TestLatency runs a single stream so we can look at
// blocking profiles for different buffer sizes.
func TestLatency(t *testing.T) {
	bufSize := 100

	fmt.Println("BufSize:", bufSize)
	stream(bufSize)
}

// TestLatencies provides a test to profile and trace channel latencies
// with a little data science sprinkled in.
func TestLatencies(t *testing.T) {
	var bufSize int
	var count int
	var first time.Duration

	pts := make(plotter.XYs, 20)

	for {

		// Perform a stream with specified buffer size.
		since := stream(bufSize)
