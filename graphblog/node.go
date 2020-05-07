// Exercise provided by Phil Pearl
// https://syslog.ravelin.com/making-something-faster-56dd6b772b83

package graphblog

import (
	"container/list"
)

type node struct {
	id  string
	adj graph
}

func (n *node) add(adjNode *node) {
	n.adj[adjNode.id] = adjNode
}

// =============================================================================

type graph map[string]*node

func newGraph() graph {
	return make(graph)
}
