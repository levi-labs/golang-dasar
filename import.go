package main

import (
	"belajar-dasar-golang/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("John Doe")
	fmt.Println(result)
	fmt.Println(helper.Application)
	checkVersion := helper.Get()
	fmt.Println(checkVersion)
}
