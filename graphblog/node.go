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

func (g graph) get(id string) *node {
	if n, found := g[id]; found {
		return n
	}

	n := node{
		id:  id,
		adj: make(graph),
	}
	g[id] = &n
	return &n
}

func (g graph) addEdge(a, b string) {
	an := g.get(a)
	bn := g.get(b)
	an.add(bn)
	bn.add(an)
}

func (g graph) diameter() int {
	var diameter int
	for id := range g {
		if df := g.longestShortestPath(id); df > diameter {
			diameter = df
		}
	}
	return diameter
}
