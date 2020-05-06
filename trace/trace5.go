// Sample program that performs a series of I/O related tasks to
// better understand tracing in Go.
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

        trace.Start(os.Stdout)
	defer trace.Stop()

	docs := make([]string, 4000)
	for i := range docs {
		docs[i] = fmt.Sprintf("newsfeed-%.4d.xml", i)
	}
