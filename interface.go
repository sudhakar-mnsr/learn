package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	worker1 := func(ctx context.Context) (interface{}, error) {
		time.Sleep(time.Millisecond)
		return "test string", nil
	}
	v1, err := retry(context.Background(), time.Second, worker1)
	fmt.Println(v1.(string), err)
