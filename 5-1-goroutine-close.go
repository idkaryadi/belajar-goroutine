package main

import (
	"fmt"
	"runtime"
)

func main() {
	// mengimplementasikan close untuk menutup channel

	runtime.GOMAXPROCS(2)
	message := make(chan string)

	// Before
	// go sendMessage(message)
	// printMessage(message)

	// After
	go sendMessageWithClose(message)
	// kalo go nya di bagian print, kena deadlock
	printMessageWithClose(message)
}

func sendMessage(ch chan string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("Data ke :: %v", i)
	}
}

func printMessage(ch chan string) {
	// disini akan timbul masalah ketika i nya lebih dari 20, yakni dead lock
	// dan kita kadang ndak pernah tahun i nya sampai berapa
	// makanya dipengiriman perlu ada close buat nutup channel
	for i := 0; i < 20; i++ {
		message := <-ch
		fmt.Println("VALUE::", message)
	}
}

// chan<- buat ngirim pesan, sama kayak sebelumnya
func sendMessageWithClose(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("Data ke:: %v", i)
	}
	// kalo gak ada close, print nya bakal kena deadlock
	close(ch)
}

// <-chan buat nerima pesan, sama kayak sebelumnya
func printMessageWithClose(ch <-chan string) {
	// value di channel bisa dilooping make for - range
	for message := range ch {
		fmt.Println("Pesan diterima::", message)
	}
}
