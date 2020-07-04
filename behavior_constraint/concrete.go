package main

import "fmt"

func main() {
	users := []user{
		{name: "Bill", email: "bill@ardanlabs.com"},
		{name: "Ale", email: "ale@whatever.com"},
	}
	s := stringifyUsers(users)
	fmt.Println(s)
