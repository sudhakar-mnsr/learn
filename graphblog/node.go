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
