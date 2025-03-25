package main

import "fmt"

// Function with named return variables
func add(a, b int) (sum int) {
    sum = a + b
    return 
}

func main() {
    result := add(3, 4)
    fmt.Println("Sum:", result) 
}
