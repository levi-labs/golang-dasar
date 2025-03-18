package main

import "fmt"

func main() {
	var (
		address     = "jalan no 11"
		phoneNumber = "08129292"
	)
	//full type
	var name string = "i use var and string as type variable"
	//use var and automate deetect type first assign
	var first_name = "first name is budig"
	//short varible declaration
	last_name := "santoso"
	fmt.Println("this is use var and string", name)

	fmt.Println("this is use var and golang automate detect type first assign is String", first_name)

	fmt.Println("this is use ", last_name)

	fmt.Println("this is use ", address+" "+phoneNumber)
}
