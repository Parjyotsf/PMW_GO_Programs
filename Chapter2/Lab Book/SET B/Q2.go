//2. WAP in go language to demonstrate use of names returns variables.
package main

import "fmt"

func calculateSumAndProduct(a, b int) (sum, product int) {
    sum = a + b
    product = a * b
    return
}

func main() {
    a, b := 5, 10
    sum, product := calculateSumAndProduct(a, b)
    fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}
