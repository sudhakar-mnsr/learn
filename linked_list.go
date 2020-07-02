package main

import "fmt"

func main() {
	var l list

	l.add("bill")
	l.add("ale")

	f := func(n *node) error {
		fmt.Println(n.Data)
		return nil
	}
	l.operate(f)
}

// =============================================================================

type node struct {
	Data string
	next *node
	prev *node
}


type list struct {
	Count int
	first *node
	last  *node
}

func (l *list) add(data string) *node {
	n := node{
		Data: data,
		prev: l.last,
	}
	l.Count++

	if l.first == nil && l.last == nil {
		l.first = &n
		l.last = &n
		return &n
	}
