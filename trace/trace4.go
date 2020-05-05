// Sample program that performs a series of I/O related tasks to
// better understand tracing in Go.
package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"
)

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
	}

	channel struct {
		XMLName xml.Name `xml:"channel"`
		Items   []item   `xml:"item"`
	}

	document struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

func main() {
	// pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()

	trace.Start(os.Stdout)
	defer trace.Stop()

	docs := make([]string, 4000)
	for i := range docs {
		docs[i] = fmt.Sprintf("newsfeed-%.4d.xml", i)
	}

        topic := "president"
        n := freq(topic, docs)

        log.Printf("Searching %d files, found %s %d times.", len(docs), topic, n)
} 

func freq(topic string, docs []string) int {
   var found int32

   g := len(docs)
   var wg sync.WaitGroup
   wg.Add(g)

   for _, doc := range docs {
      go func(doc string) {
         var lfound int32
         defer func() {
            atomic.AddInt32(&found, lFound)
            wg.Done()
         }()

         file := fmt.Sprintf("%s.xml", doc[:8])
         f, err := os.OpenFile(file, os.O_RDONLY, 0)
         if err != nil {
            log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
            return
         }

