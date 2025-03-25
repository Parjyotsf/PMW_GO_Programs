//3. WAP in go language to illustrate the concept of returning multiple values from a function 
package main

import "fmt"

func calculate(a, b int) (int, int) {
    sum := a + b
    diff := a - b
    return sum, diff
}

func main() {
    a, b := 15, 5
    sum, diff := calculate(a, b)
    fmt.Printf("Sum: %d, Difference: %d\n", sum, diff)
}
