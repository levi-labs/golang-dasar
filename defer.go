package main

import "fmt"

func logging() {
	fmt.Println("start logging")
}

func errorss(error bool) {
	defer logging()
	if error {
		fmt.Println("error")
	}
}
func main() {
	fmt.Println("start")
	errorss(true)
}
