//2. WAP in go language to print recursive sum of digits of given number. 
package main

import "fmt"

func sumOfDigits(n int) int {
    if n == 0 {
        return 0
    }
    return n%10 + sumOfDigits(n/10) 
}

func main() {
    var num int

    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num < 0 {
        num = -num
    }

    result := sumOfDigits(num)

    fmt.Printf("The recursive sum of digits of %d is: %d\n", num, result)
}
