package main

import(
	"fmt"
)

func main(){
	sumFirst, sumSecond := 0, 0
	for i := 1; i <= 100; i++{
		sumFirst += i
		sumSecond += i * i
	}
	finalSumFirst := sumFirst * sumFirst
	result := finalSumFirst - sumSecond
	fmt.Println(result)
}

/* alternative method
func main() {
    result := ((100*101)/2)*((100*101)/2) - ((100 * 101 * 201) / 6)
    println(result)
}

*/