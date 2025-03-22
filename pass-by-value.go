package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"jakarta", "subang", "yogyakarta"}
	address2 := address1

	address2.City = "bandung"

	fmt.Println(address1)
	fmt.Println(address2)
}
