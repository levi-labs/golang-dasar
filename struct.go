package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//Struct Literal

func (c Customer) sayHello(name string) {
	fmt.Println("Hello", c.Name, " My name is ", name)
}

func main() {
	//if struct belum di definisikan valunya maka ia akan menggunakan default value nya
	var customer Customer

	customer.Name = "John Doe"
	customer.Address = "123 street 123"
	customer.Age = 12

	fmt.Println(customer.Name)
	john := Customer{
		Name:    "John",
		Address: "Jane Doe",
		Age:     12,
	}
	fmt.Println(john)

	//OR

	james := Customer{"james", "street number 212", 21}
	fmt.Println(james)

	//Method Struct
	james.sayHello("Doni")
}
