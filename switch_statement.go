package main

import "fmt"

func main() {
	name := "Joko Anwar"

	switch name {
	case "Joko Anwar":
		fmt.Println("The name is Joko Anwar")
	case "Sanusi ahmad":
		fmt.Println("The name is Sanusi ahmad")
	default:
		fmt.Println("Please enter a name")
	}

	//switch with short statement
	switch length := len(name); length > 0 {
	case true:
		fmt.Println("The name is ", name)
	case false:
		fmt.Println("The name is ", name)
	}

	//switch use case condition
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("The name is too long")
	case length < 3:
		fmt.Println("The name is too short")
	}

}
