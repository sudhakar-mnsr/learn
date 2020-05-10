// This sample program demonstrates how to create a simple chat system.
package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/ardanlabs/gotraining/topics/go/concurrency/patterns/chat"
)

func main() {
	cr := chat.New()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting Down Started")
	cr.Close()
	log.Println("Shutting Down Completed")
}
