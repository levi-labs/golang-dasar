package main

import "fmt"

func getSumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	total := getSumAll(1, 2, 3, 4, 5)
	fmt.Println(total)

	//Slice parameter
	numbers := []int{1, 2, 3, 4, 7}
	fmt.Println(getSumAll(numbers...))
}
