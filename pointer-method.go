package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) isMarried() {
	man.Name = "Mr. " + man.Name
}
func main() {
	budi := Man{"budi"}
	budi.isMarried()
	fmt.Println(budi.Name)
}
