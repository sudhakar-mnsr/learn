  
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	worker1 := func(ctx context.Context) (string, error) {
		time.Sleep(time.Millisecond)
		return "test string", nil
	}
	v1, err := retry(string)(context.Background(), time.Second, worker1)
	fmt.Println(v1, err)

	worker2 := func(ctx context.Context) (int, error) {
		time.Sleep(time.Millisecond)
		return 9999999, nil
	}
	v2, err := retry(int)(context.Background(), time.Second, worker2)
	fmt.Println(v2, err)

	worker3 := func(ctx context.Context) (*user, error) {
		time.Sleep(time.Millisecond)
		return &user{"bill", "b@email.com"}, nil
	}
	v3, err := retry(*user)(context.Background(), time.Second, worker3)
	fmt.Println(v3, err)
}
