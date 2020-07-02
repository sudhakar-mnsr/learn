package main

import "fmt"

func main() {
	var l list(string)

	l.add("bill")
	l.add("ale")

	f := func(n *node(string)) error {
		fmt.Println(n.Data)
		return nil
	}
	l.operate(f)
}

// =============================================================================

type scalarOnly interface {
	type int, int8, int16, int32, int64, string
}

type node(type T scalarOnly) struct {
	Data T
	next *node(T)
	prev *node(T)
}

type list(type T scalarOnly) struct {
	Count int
	first *node(T)
	last  *node(T)
}

func (l *list(T)) add(data T) *node(T) {
	n := node(T) {
		Data: data,
		prev: l.last,
	}
	l.Count++
