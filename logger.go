// This sample program demonstrates how the logger package works.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

// device allows us to mock a device we write logs to.
type device struct {
	mu      sync.RWMutex
	problem bool
}
