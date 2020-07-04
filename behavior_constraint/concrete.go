package main

import "fmt"

func main() {
	users := []user{
		{name: "Bill", email: "bill@ardanlabs.com"},
		{name: "Ale", email: "ale@whatever.com"},
	}
	s := stringifyUsers(users)
	fmt.Println(s)

	customers := []customer{
		{name: "Google", email: "you@google.com"},
		{name: "MSFT", email: "you@msft.com"},
	}

	s = stringifyCustomers(customers)
	fmt.Println(s)
}

// =============================================================================

func stringifyUsers(users []user) []string {
	ret := make([]string, 0, len(users))
	for _, user := range users {
		ret = append(ret, user.String())
	}
	return ret
}

func stringifyCustomers(customers []customer) []string {
	ret := make([]string, 0, len(customers))
	for _, customer := range customers {
