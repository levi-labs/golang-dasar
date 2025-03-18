package main

import "fmt"

func main() {
	name := [...]string{"apple", "banana", "grape", "eggplant", "avocado", "pineapple", "strawberry", "jackfruit"}
	//jadi slice itu bisa dibuat dari array yang udah ada,
	//indikator pertama itu adalah index ( 4 )  dan indikator ke 2 adalah range nya dalam hal ini 6
	//artinya akan di ambil data dari index ke 4 sampai sebelum index ke 6 [ avocado dan pinepple]

	/*
			function yang bisa digunakan pada slice
			len()
		 	cap(slince) untuk mendapatkan slice
			append(slice,data) membuat slice baru dan menambahkan data baru di akhir slice, jika kapasistas penuh
			maka akan dibuatkan array baru
			make([]TypeData,length,capacity) membuat slice baru
			copy(destination,source) menyalin source ke destination
	*/

	slice1 := name[4:6]

	fmt.Println(slice1)

	//dar index ke 4
	slice2 := name[4:]
	fmt.Println(slice2)

	//sebelum index ke 4
	slice3 := name[:4]
	fmt.Println(slice3)

	//ambil semuanya
	slice4 := name[:]
	fmt.Println(slice4)

	//make slice manual

	var slice5 []string = name[:]

	fmt.Println(slice5)

	// append akan menghasilkan array baru

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daysSlice1 := days[5:]
	daysSlice1[0] = "sabtu baru"
	daysSlice1[1] = "minggu baru"

	fmt.Println(days)

	daySlice2 := append(daysSlice1, "Libur Baru")
	fmt.Println(daySlice2)
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	//make slice from sratch

	newSlice := make([]int, 3, 2)
	newSlice[0] = 1
	newSlice[1] = 2
	newSlice[2] = 3
	fmt.Println(newSlice)

	newSlice2 := append(newSlice, 7)
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	//[...]int{1,2,3} array
	//[]int{1,2,3} slice
}
