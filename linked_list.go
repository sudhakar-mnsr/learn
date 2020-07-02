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
