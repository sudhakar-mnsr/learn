// Exercise provided by Phil Pearl
// https://syslog.ravelin.com/making-something-faster-56dd6b772b83

package graphblog

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

type edge struct {
	a string
	b string
}

// =============================================================================

type edges []edge

func (e edges) build(g graph) {
	for _, edge := range e {
		g.addEdge(edge.a, edge.b)
	}
}
