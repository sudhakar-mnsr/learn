// Sample program that performs a series of I/O related tasks to 
// better understand tracing in Go
package main

import (
   "context"
   "encoding/xml"
   "fmt"
   "io/ioutil"
   "log"
   "os"
   "runtime"
   "runtime/pprof"
   "runtime/trace"
   "strings"
   "sync"
   "sync/atomic"
)

type (
   item struct {
      XMLName xml.Name `xml:"item"`
      Title   string `xml:"title"`
      Description string `xml: "description"`
   }

   channel struct {
      XMLName xml.Name `xml:"channel"`
      Items []item `xml:"item"`
   }

   document struct {
      XMLName xml.Name `xml:"rss"`
      Channel channel `xml:"channel"`
   }
)

func main() {
   pprof.StartCPUProfile(os.Stdout)
   defer pprof.StopCPUProfile()

   docs := make([]string, 4000)
   for i := range docs {
      docs[i] = fmt.Sprintf("newsfeed-%.4d.xml", i)
   }

   topic := "president"
   n := freq(topic, docs)
   // n := freqConcurrent(topic, docs)
   // n := freqConcurrentSem(topic, docs)
   // n := freq(topic, docs)
   // n := freqProcessorsTasks(topic, docs)
   // n := freqActor(topic, docs)

   log.Printf("Searching %d files, found %s %d times.", len(docs), topic, n)
}

func freq(topic string, docs []string) int {
   var found int

   for _, doc := range docs {
      file := fmt.Sprintf("%s.xml", doc[:8])
      f, err := os.OpenFile(file, os.O_RDONLY, 0)
      if err != nil {
         log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
         return 0
      }

      data, err := ioutil.ReadAll(f)
      if err != nil {
         f.Close()
         log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
         return 0
      }
      f.Close()

      var d document
      if err := xml.Unmarshal(data, &d); err != nil {
         log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
         return 0
      }

      for _, item := range d.Channel.Items {
         if strings.Contains(item.Title, topic) {
            found++
            continue
         }
         
         if strings.Contains(item.Description, topic) {
            found++
         }
      }
   }
   return found
}
