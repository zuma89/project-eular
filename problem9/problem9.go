package main

import (
	"fmt"
	"math"
)

func isSquare(num int) (bool, int) {
	sqrt := int(math.Sqrt(float64(num)))
	if sqrt*sqrt == num {
		return true, sqrt
	}
	return false, -1
}

func main() {
	for a := 1; a < 500; a++ {
		for b := a; b < 500; b++ {
			result, c := isSquare(a*a + b*b)
			if result && a+b+c == 1000 {
				fmt.Println(a * b * c)
				break loop
			}
		}
	}
}

/*
def prod_triplet_w_sum(n):
    for i in range(1,n,1):
        for j in range(1,n-i,1):
            k = n-i-j
            if i**2+j**2==k**2:
                return i*j*k
    return 0
*/
