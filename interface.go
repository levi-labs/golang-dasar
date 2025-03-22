package main

import "fmt"

type Animal interface {
	getKingdom() string
}

type Dog struct {
	Name    string
	Kingdom string
}

func sayHello(a Animal) {
	fmt.Println(a.getKingdom())
}

func (d *Dog) getKingdom() string {
	d.Kingdom = "Karnivora"
	return d.Kingdom
}
func main() {

	silvester := Dog{
		Name: "Silvester",
	}
	fmt.Println(silvester.getKingdom())
	fmt.Println(sayHello(silvester))

}
