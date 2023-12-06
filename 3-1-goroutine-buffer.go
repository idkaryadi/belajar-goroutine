package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	message := make(chan string, 3)

	go kirimValue(message)

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		message <- fmt.Sprintf("angka ke %d", i)
	}
}

func kirimValue(ch chan string) {
	for {
		// kalo value diatas for, jadi infinite loop
		value := <-ch
		fmt.Println("terima data", value)
	}
}
