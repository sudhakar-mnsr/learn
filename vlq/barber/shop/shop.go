// Package shop implements the sleeping barber problem.
// There is one barber in the barber shop, one barber chair and n chairs for
// waiting customers. If there are no customers, the barber sits down in the
// barber chair and takes a nap. An arriving customer must wake the barber.
// Subsequent arriving customers take a waiting chair if any are empty or
// leave if all chairs are full.
//
// Have the ability to close the shop even if new customers are entering.
// Customers looking for a chair should run on their own goroutine.
//
// Task: Change EnterCustomer so a customer can wait for a specified amount
// of time for a chair to open up.

package shop

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// ErrShopClosed is returned when the shop is closed.
	ErrShopClosed = errors.New("shop closed")

	// ErrNoChair is returned when all the chairs are occupied.
	ErrNoChair = errors.New("no chair available")
)

// customer represents a customer to be serviced.
type customer struct {
	name string
}

// Shop represents the barber's shop which contains chairs for customers
// that customers can occupy and the barber can service. The shop can
// be closed for business.
type Shop struct {
	open    int32          // Determines if the shop is open for business.
	chairs  chan customer  // The set of chairs in the shop.
	wgClose sync.WaitGroup // Provides support for closing the shop.
	wgEnter sync.WaitGroup // Tracks customers entering the shop.
}

// Open creates a new shop for business and gets the barber working.
func Open(maxChairs int) *Shop {
	s := Shop{
		chairs: make(chan customer, maxChairs),
	}
	atomic.StoreInt32(&s.open, 1)

	// Get the barber working.
	s.wgClose.Add(1)
	go func() {
		defer s.wgClose.Done()
		for cust := range s.chairs {
			fmt.Printf("Barber servicing customer %q\n", cust.name)
