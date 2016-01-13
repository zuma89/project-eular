package main

import "fmt"

func main(){
	fib1, fib2, sum := 1, 2, 0
	for fib2 < 4000000{
		if fib2 % 2 == 0{
			sum += fib2
		}
		fib1, fib2 = fib2, fib1 + fib2
	}
	fmt.Println(sum)
}