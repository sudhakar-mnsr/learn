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

