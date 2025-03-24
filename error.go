package main

import (
	"errors"
	"fmt"
)

func Divide(value int, divided int) (int, error) {
	if divided == 0 {
		return 0, errors.New("division by zero")
	} else {
		return value / divided, nil
	}
}

func main() {
	result, err := Divide(10, 0)
	if err == nil {
		fmt.Println("result:", result)
	} else {
		fmt.Println("error:", err.Error())
	}
}
