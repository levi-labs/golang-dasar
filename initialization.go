package main

import (
	"belajar-dasar-golang/database"
	_ "belajar-dasar-golang/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetConnection())
}
