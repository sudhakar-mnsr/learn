// Example provided with help from Fatih Arslan and Gabriel Aszalos.

// Package pool manages a user defined set of resources.
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)
