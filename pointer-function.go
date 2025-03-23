package main

import "fmt"

func changeAddress(address *Address) {
	address.Country = "New York"
}

type Address struct {
	City, Province, Country string
}

func main() {
	address := Address{"Surabaya", "Jakarta Timur", ""}
	changeAddress(&address)
	fmt.Println(address)
}
