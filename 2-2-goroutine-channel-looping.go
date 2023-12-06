package main

import "fmt"

func main() {
	message := make(chan string)

	for _, value := range []string{"hallo", "haii", "selamat"} {
		go func(word string) {
			newWord := fmt.Sprintf("kata:: %s", word)
			message <- newWord
		}(value)
	}

	// 3 dari panjang kata, kalo kurang dari 3 ndakpapa walaupun ada kata yg tidak diterima
	// kalo lebih dari 3 bakal error karena ndak ada yg diterima => goroutines are asleep - deadlock!
	for i := 0; i < 2; i++ {
		fmt.Println("diterima", <-message)
	}
}
