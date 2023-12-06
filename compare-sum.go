package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	var randomNum []int
	for i := 0; i < 100; i++ {
		r := rand.Intn(100)
		randomNum = append(randomNum, r)
	}

	fmt.Println("res", randomNum)

	start := time.Now()
	// fmt.Println("strat", start)
	var sum int
	for _, v := range randomNum {
		sum += v
	}
	fmt.Println("res", sum)
	end := time.Now()
	// fmt.Println("end", end)
	fmt.Println("duration", end.Sub(start).Nanoseconds())

	start2 := time.Now()
	runtime.GOMAXPROCS(1)
	message := make(chan int)
	for i := 0; i < len(randomNum); i = i + 25 {
		// fmt.Println("i", i)
		// fmt.Println("i+25", i+25)
		// fmt.Println("len", len(randomNum[i:i+25]))
		go total(message, randomNum[i:i+25])
	}

	sum2 := 0
	for i := 0; i < 4; i++ {
		res := <-message
		sum2 += res
	}
	fmt.Println("sum2", sum2)
	end2 := time.Now()
	fmt.Println("duration 2", end2.Sub(start2).Nanoseconds())
}

func total(ch chan int, randomNum []int) {
	var sum int
	for _, v := range randomNum {
		sum += v
	}
	fmt.Println("sum", sum)
	ch <- sum
}
