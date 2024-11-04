/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10001st prime number?
*/

package main

import(
	"fmt"
)

func main(){
	arr := make([]bool, 105000)
    arr[0], arr[1] = true, true
    count, prime := 2, 3
    var i int
    for {
        for i = 2 * prime; i < len(arr); i += prime {
            arr[i] = true
        }
        for i = prime + 2; i < len(arr) && arr[i]; i += 2 {
        }
        if i < len(arr) {
            prime = i
            count++
            if count == 10001 {
                fmt.Println(prime)
                break
            }
        } else {
            break
        }
    }
}