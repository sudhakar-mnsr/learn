// Sample program that performs a series of I/O related tasks to
// better understand tracing in Go.
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

