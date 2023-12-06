package main

import "fmt"

func main() {
	// tidak akan memunculkan sesuatu karena main selesai setelah memanggil printAngka
	go printAngka("lima")

	// var hehe string
	// fmt.Scanln(&hehe)
	// fmt.Println("hehhe", hehe)
}

func printAngka(angka string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "muncul", angka)
	}
}
