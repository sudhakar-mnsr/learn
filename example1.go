// Using the template, declare a set of concrete types that implement the set
// of predefined interface types. Then create values of these types and use
// them to complete a set of predefined tasks.
package main

// Add import(s).
import "fmt"

// administrator represents a person or other entity capable of administering
// hardware and software infrastructure.
type administrator interface {
	administrate(system string)
}

// developer represents a person or other entity capable of writing software.
type developer interface {
	develop(system string)
}

// =============================================================================

// adminlist represents a group of administrators.
type adminlist struct {
	list []administrator
}

// Enqueue adds an administrator to the adminlist.
func (l *adminlist) Enqueue(a administrator) {
	l.list = append(l.list, a)
}
