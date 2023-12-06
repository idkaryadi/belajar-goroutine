package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	// 3 merupakan ukurannya, jadi valuenya ada 3+1, karena berawal dari 0
	message := make(chan int, 3)

	go func() {
		for {
			i := <-message
			fmt.Println("nerima pesan", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 15; i++ {
		// kalo pesannya sudah 4 bakal ngirim semua datanya
		// bakal nambah lagi ketika datanya sudah terkirim dan ada yang nerima
		fmt.Println("kirim value", i)
		message <- i
	}

	time.Sleep(10 * time.Second)
}
