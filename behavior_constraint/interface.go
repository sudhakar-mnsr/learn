package main

import "fmt"

func main() {
	users := []user{
		{name: "Bill", email: "bill@ardanlabs.com"},
		{name: "Ale", email: "ale@whatever.com"},
	}

	s := stringify(users)
	fmt.Println(s)

	customers := []customer{
		{name: "Google", email: "you@google.com"},
		{name: "MSFT", email: "you@msft.com"},
	}
	s = stringify(customers)
	fmt.Println(s)
}

// =============================================================================

func stringify(v interface{}) []string {
	switch list := v.(type) {
	case []user:
		ret := make([]string, 0, len(list))
		for _, value := range list {
			ret = append(ret, value.String())
		}
		return ret
	case []customer:
		ret := make([]string, 0, len(list))
		for _, value := range list {
			ret = append(ret, value.String())
		}
		return ret
	}
	return nil
}

// =============================================================================

type user struct {
	name  string
	email string
}

func (u user) String() string {
	return fmt.Sprintf("{type: \"user\", name: %q, email: %q}", u.name, u.email)
}
