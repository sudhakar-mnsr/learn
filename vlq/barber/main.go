// Program creates customers for the simulation of the sleeping barber problem
// implemented in the shop package.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/fun/barber/shop"
)

func main() {
	const maxChairs = 10
	s := shop.Open(maxChairs)

	// Create a goroutine than is constantly, but inconsistently, generating
	// customers who are entering the shop.
	go func() {
