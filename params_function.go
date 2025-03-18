package main

import "fmt"

type Func func(string) string

func sayHelloWithFilter(name string, filter Func) {
	filteredName := filter(name)
	fmt.Println(filteredName)
}

func spamFilter(name string) string {
	if name == "asu" {
		sensor := ""
		for i := 0; i < len(name); i++ {
			//make sensor * with length
			sensor += "*"
		}
		return sensor
	}
	return name
}
func main() {
	sayHelloWithFilter("asu", spamFilter)

}
