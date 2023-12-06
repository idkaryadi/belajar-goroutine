package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(5)

	// goroutine adalah light weight thread, karena gak memakan banyak memori
	// di bahasa pemrograman lain ada thread namun memakan banyak memori
	// penggunaan goroutine go namaFunc()
	go printAngka("satu")
	printAngka("dua")

	// untuk melakukan blocking pada go routine, kalo tidak ada, printAngka satu tidak akan dieksekusi
	// karena program sudah berakhir
	fmt.Scanln()
}

func printAngka(angka string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "muncul", angka)
	}
}
