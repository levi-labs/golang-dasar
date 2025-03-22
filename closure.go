package main

import "fmt"

func main() {
	count := 0

	increament := func() {
		fmt.Println(count)
		count++
	}

	increament()
	increament()
	increament()

	//output 3 , karena increament dapat mengakses nilai terakhir pada count
	fmt.Println("output", count)
}
