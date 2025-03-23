package main

import "fmt"

type Addresss struct {
	City, Province, Country string
}

func main() {
	var address1 Addresss = Addresss{"jakarta", "subang", "yogyakarta"}
	var address2 *Addresss = &address1

	address2.City = "bandungs"

	fmt.Println(address1) //akan b erubah
	fmt.Println(address2)
}
