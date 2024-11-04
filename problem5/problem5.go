/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import "fmt"

func main(){
	divisor := []int{2,3,5,7,11,13,17,19} 
	result := 1
	var num int
	for i := range divisor{
		num = divisor[i]
		for num <= 20{
			num *= divisor[i]
		}
		num /= divisor[i]
		result *= num
	}
	fmt.Println(result)
}

