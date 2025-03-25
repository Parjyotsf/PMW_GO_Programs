package main

import "fmt"

// Factorial function using direct recursion
func factorial(n int) int {
    if n == 0 {
        return 1 // Base case
    }
    return n * factorial(n-1) // Recursive call
}

func main() {
    num := 5
    fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}
