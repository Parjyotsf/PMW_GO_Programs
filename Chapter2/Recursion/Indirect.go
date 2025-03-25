package main

import "fmt"

// Indirect Recursion: function odd calls even, and even calls odd
func odd(n int) {
    if n <= 0 {
        return
    }
    fmt.Println(n)
    even(n - 1) // Calling even function
}

func even(n int) {
    if n <= 0 {
        return
    }
    fmt.Println(n)
    odd(n - 1) // Calling odd function
}

func main() {
    odd(5) // Output: 5, 4, 3, 2, 1
}
