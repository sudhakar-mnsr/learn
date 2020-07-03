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
