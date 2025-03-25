package main

import "fmt"

// Factorial function using recursion
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1) // Recursive call
}

func main() {
    result := factorial(5)
    fmt.Println("Factorial of 5:", result) // Output: 120
}
