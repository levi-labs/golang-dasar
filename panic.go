package main

import "fmt"

func endApp() {
	fmt.Println("endApp")
}
func runApp(error bool) {
	defer endApp() //tetap dijalankan meskipun error
	if error {
		panic("error")
	}
}

func main() {
	runApp(true)
}
