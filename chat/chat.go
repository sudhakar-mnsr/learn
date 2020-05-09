// Package chat implements a basic chat room.
package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

// temporary is declared to test for the existence of the method coming
// from the net package.
type temporary interface {
	Temporary() bool
}

// message is the data received and sent to users in the chatroom.
type message struct {
	data string
	conn net.Conn
}

// client represents a single connection in the room.
type client struct {
	name   string
	room   *Room
	reader *bufio.Reader
	writer *bufio.Writer
	wg     sync.WaitGroup
	conn   net.Conn
}

// read waits for message and sends it to the chatroom for processing.
func (c *client) read() {
	for {

		// Wait for a message to arrive.
		line, err := c.reader.ReadString('\n')

		if err == nil {
			c.room.outgoing <- message{
				data: line,
				conn: c.conn,
			}
			continue
		}


		if e, is := err.(temporary); is && !e.Temporary() {
			log.Println("Temporary: Client leaving chat")
			c.wg.Done()
			return
		}

		if err == io.EOF {
			log.Println("EOF: Client leaving chat")
			c.wg.Done()
			return
		}

		log.Println("read-routine", err)
	}
}


// write is a goroutine to handle processing outgoing
// messages to this client.
func (c *client) write(m message) {
	msg := fmt.Sprintf("%s %s", c.name, m.data)
	log.Printf(msg)

	c.writer.WriteString(msg)
	c.writer.Flush()
}
