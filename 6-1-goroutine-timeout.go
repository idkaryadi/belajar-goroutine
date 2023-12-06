package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// menutup channel ketika sudah waktu tertentu
func main() {
	runtime.GOMAXPROCS(2)

	message := make(chan int)
	go sendData(message)
	receiveData(message)
}

func sendData(ch chan int) {
	for i := 0; true; i++ {
		fmt.Println("i", i)
		ch <- i
		waktuSleep := rand.Intn(6)
		fmt.Println("SLEEP", waktuSleep)
		time.Sleep(time.Duration(waktuSleep) * time.Second)
	}
}

func receiveData(ch chan int) {
	// kalo gak make loop, bakal tetep looping, walaupun udah di break
loop:
	for {
		select {
		case data := <-ch:
			fmt.Println("nerima data::", data)
			// fmt.Sprintf("nerima data:: %v", data)
		case <-time.After(5 * time.Second):
			fmt.Println("close after 5 second")
			break loop
		}
	}

}
