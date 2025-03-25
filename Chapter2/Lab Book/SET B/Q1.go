//1. WAP in go language to swap two numbers using call by reference concept. 
package main

import "fmt"

func swap(a *int, b *int) {
    *a, *b = *b, *a
}

func main() {
    var num1, num2 int

    fmt.Print("Enter the first number: ")
    fmt.Scan(&num1)
    fmt.Print("Enter the second number: ")
    fmt.Scan(&num2)

    fmt.Printf("Before Swap: num1 = %d, num2 = %d\n", num1, num2)

   
    swap(&num1, &num2)

    fmt.Printf("After Swap: num1 = %d, num2 = %d\n", num1, num2)
}
