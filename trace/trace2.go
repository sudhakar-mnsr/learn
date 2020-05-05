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
      XMLName xml.Name `xml:"item"`
      Title string `xml:"title"`
      Description string `xml:"description"`
   }
   channel struct {
      XMLName xml.Name `xml."channel"`
      Items []item `xml:"item"`
   }
   document struct {
      XMLName xml.Name `xml:"rss"`
      Channel channel `xml:"channel"`
   }
)

func main() {
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

func freq(topic string, docs[]string) int {
   var found int32
   
   g := len(docs)
   var wg sync.WaitGroup
   wg.Add(g)
   
   for _, doc := range docs {
      go func(doc string) {
         defer wg.Done()
         file := fmt.Sprintf("%s.xml", doc[:8])
         f, err := os.OpenFile(file, os.O_RDONLY, 0)
         if err != nil {
                 log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
                 return
         }
   
         data, err := ioutil.ReadAll(f)
         if err != nil {
            f.Close()
            log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
            return
         }
         f.Close()
   
         var d document
   
         if err := xml.Unmarshal(data, &d); err != nil {
            log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
            return
         }
   
         for _, item := range d.Channel.Items {
            if strings.Contains(item.Title, topic) {
               atomic.AddInt32(&found, 1)
               continue
            }
            
            if strings.Contains(item.Description, topic) {
               atomic.AddInt32(&found, 1)
            }
         }
      }(doc)
   }
   wg.Wait() 
   return int(found)
}
