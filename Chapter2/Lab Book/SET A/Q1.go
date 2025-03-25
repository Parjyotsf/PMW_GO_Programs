//1. WAP in go language to print addition of two number using function. 
package main

import "fmt"

func addNumbers(a int, b int) int {
    return a + b
}

func main() {
    var num1, num2 int

    
    fmt.Print("Enter the first number: ")
    fmt.Scan(&num1)
    fmt.Print("Enter the second number: ")
    fmt.Scan(&num2)

    sum := addNumbers(num1, num2)

  
    fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)
}
