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

// =============================================================================

func TestDiameter(t *testing.T) {
	tests := []struct {
		name        string
		edges       edges
		expDiameter int
	}{
		{
			name: "empty",
		},

		{
			name:        "1edge",
			edges:       edges{{"a", "b"}},
			expDiameter: 1,
		},
		{
			name:        "3inline",
			edges:       edges{{"a", "b"}, {"b", "c"}},
			expDiameter: 2,
		},
		{
			name:        "4inline",
			edges:       edges{{"a", "b"}, {"b", "c"}, {"c", "d"}},
			expDiameter: 3,
		},
		{
			name:        "triangle",
			edges:       edges{{"a", "b"}, {"b", "c"}, {"a", "c"}},
			expDiameter: 1,
		},
