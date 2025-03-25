//1. WAP in go language to illustrate the concept of call by value.
package main

import "fmt"

func modifyValue(val int) {
    val = 20
}

func main() {
    num := 10
    fmt.Printf("Before function call: %d\n", num)
    modifyValue(num)
    fmt.Printf("After function call: %d\n", num)
}
