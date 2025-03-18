package main

import "fmt"

func recursiveLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= 1
	}
	return result
}
func recursiveFunction(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFunction(value-1)
	}
}

func main() {
	recursiveManual := 5 * 4 * 3 * 2 * 1
	fmt.Println(recursiveManual)

	fmt.Println(recursiveFunction(5))
}
