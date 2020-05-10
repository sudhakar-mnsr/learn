// This sample program demonstrates how to use the pool package
// to share a simulated set of database connections.
package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sudhakarmnsr/learn/pool"
)

const (
	maxGoroutines = 25 // the number of routines to use.
	numPooled     = 2  // number of resources in the pool
)

// dbConnection simulates a resource to share.
type dbConnection struct {
	ID int32
}

// Close implements the io.Closer interface so dbConnection
// can be managed by the pool. Close performs any resource
// release management.
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter provides support for giving each connection a unique id.
var idCounter int32

// createConnection is a factory method that will be called by
// the pool when a new connection is needed.
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

// performQueries tests the resource pool of connections.
func performQueries(query int, p *pool.Pool) {

	// Acquire a connection from the pool.
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// Release the connection back to the pool.
	defer p.Release(conn)

	// Wait to simulate a query response.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("Query: QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// Create the pool to manage our connections.
	p, err := pool.New(numPooled, createConnection)
	if err != nil {
		log.Println(err)
		return
	}
