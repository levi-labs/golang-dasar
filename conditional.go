package main

import "fmt"

func main() {

	name := "budi"

	if name == "budi" {
		fmt.Println("Hello" + " " + name)
	} else {
		fmt.Println("Hello what is your name" + "!")
	}

	//if short statment
	if length := len(name); length > 0 {
		fmt.Println("The name is ", name+" and the length is ", length)
	}
}
