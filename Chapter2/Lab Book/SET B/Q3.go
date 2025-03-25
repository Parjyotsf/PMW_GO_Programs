//3. WAP in go language to show the compiler throws an error if a variable is declared but not used.
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
