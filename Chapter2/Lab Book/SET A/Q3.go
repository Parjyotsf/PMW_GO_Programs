//3. WAP in go language using function to check whether accepts number is palindrome or not.
package main

import "fmt"

func isPalindrome(num int) bool {
    original := num
    reversed := 0

    for num > 0 {
        digit := num % 10
        reversed = reversed*10 + digit
        num /= 10
    }

    return original == reversed
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num < 0 {
        fmt.Println("Negative numbers cannot be palindromes.")
        return
    }

    if isPalindrome(num) {
        fmt.Printf("%d is a palindrome.\n", num)
    } else {
        fmt.Printf("%d is not a palindrome.\n", num)
    }
}
