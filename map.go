package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Kameroon"
	//person["age"] = "23"

	person := map[string]string{
		"name": "John",
		"age":  "25",
	}
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person)

	book := make(map[string]string)
	book["name"] = "Kameroon"
	book["age"] = "25"
	book["gender"] = "male"

	fmt.Println(book)

	delete(book, "age")
	fmt.Println(book)
}
