package main

import "fmt"

func main() {
    fmt.Println(Add(10, 20))
    fmt.Println(Add("A", "B"))
    // fmt.Println(Add(3.14159, 2.96))

    d := []duration{5000, 10, 40}
    fmt.Println("Matched at index: ", match(d, 10))

    // f := []fruit{"apple", "orange", "banna"}
    // fmt.Println("Matched at index: ", match(f, "banna"))
}

// =============================================================================

type addOnly interface {
	type string, int, int8, int16, int32, int64
}

func Add(type T addOnly)(v1 T, v2 T) T {
    return v1 + v2
}

// =============================================================================

type matcher(type T) interface {
    type int, int8, int16, int32, int64, float32, float64
    match(find T) bool
}
