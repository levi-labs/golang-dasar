package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//var address1 Address = Address{"jakarta", "subang", "yogyakarta"}
	//var address2 *Address = &address1
	//
	//address2.City = "bandung"

	//address2 = &Address{"Jakarta", "DKI JAKARTA", "INDONESIA"} //berubah address 2 namun address 1 tetap bandung

	var address1 Address = Address{"jakarta", "subang", "yogyakarta"}
	var address2 *Address = &address1

	address2.City = "bandung"

	*address2 = Address{"Surabaya", "Jakarta Timur", "INDONESIA"} // merubah semua nilai termasuk semua objek yg mengacu pada address 1
	fmt.Println(address1)                                         //akan b erubah
	fmt.Println(address2)
	fmt.Println(address1)
}
