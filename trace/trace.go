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

func freqConcurrent(topic string, docs []string) int {
	var found int32

	g := len(docs)
	var wg sync.WaitGroup
	wg.Add(g)

	for _, doc := range docs {
		go func(doc string) {
			var lFound int32
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
					lFound++
					continue
				}

				if strings.Contains(item.Description, topic) {
					lFound++
				}
			}
		}(doc)
	}

	wg.Wait()
	return int(found)
}

func freqConcurrentSem(topic string, docs []string) int {
	var found int32

	g := len(docs)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan bool, runtime.GOMAXPROCS(0))

	for _, doc := range docs {
		go func(doc string) {
			ch <- true
			{
				var lFound int32
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
						lFound++
						continue
					}

					if strings.Contains(item.Description, topic) {
						lFound++
					}
				}
			}
			<-ch
		}(doc)
	}

	wg.Wait()
	return int(found)
}

func freqProcessors(topic string, docs []string) int {
	var found int32

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for i := 0; i < g; i++ {
		go func() {
			var lFound int32
			defer func() {
				atomic.AddInt32(&found, lFound)
				wg.Done()
			}()

			for doc := range ch {
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
						lFound++
						continue
					}

					if strings.Contains(item.Description, topic) {
						lFound++
					}
				}
			}
		}()
	}

	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	wg.Wait()
	return int(found)
}
