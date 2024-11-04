/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import(
	"fmt"
	"strconv"
)

func isPalindrome(input string) bool {
	i := 0
    for i < len(input)/2{
        if input[i] != input[len(input)-1-i] {
            return false
        }
		i++
    }
    return true
}

func main() {
    max, multiply := 0, 0
    for i := 100; i <= 999; i++ {
        for j := i; j <= 999; j++ {
            multiply = i * j
            if isPalindrome(strconv.Itoa(multiply)) && multiply > max {
                max = multiply
            }
        }
    }
    fmt.Println(max)
}
