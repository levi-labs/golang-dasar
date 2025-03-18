package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	//counter with statement

	for i := 1; i <= 10; i++ {
		fmt.Println("Iteration", i)
	}

	//for range
	names := []string{"John", "Paul", "George", "Ringo"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	//jika yang dibutuhkan hanya sebuah value, maka index{key) diubah menjadi _ agar tidak terjadi error
	for index, name := range names {
		fmt.Println(index, name)
	}

}
