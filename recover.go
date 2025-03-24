package main

import "fmt"

func logg() {
	fmt.Println("start logging")
	message := recover()
	fmt.Println("ini Error", message)
}
func errorrr(error bool) {
	defer logg()
	if error {
		panic("Something went wrong")
	}
}
func main() {
	errorrr(true)
}
