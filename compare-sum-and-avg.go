package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// kalo pas pertama kali dijalankan menang yg go routine, setelahnya menang yg biasa
// kalo go routinenya mau menang, bisa nama gomaxprocs, tapi bedanya bakal dikit
func main() {
	// gak make goroutine
	var numbers []int
	for i := 0; i < 10000000; i++ {
		r := rand.Intn(10000)
		numbers = append(numbers, r)
	}

	// fmt.Println("res", numbers)

	start1 := time.Now()
	max1 := getMaxArr(numbers)
	fmt.Println("max1", max1)

	avg1 := getAverageArr(numbers)
	fmt.Println("avg1", avg1)

	end1 := time.Now()
	fmt.Println("start1", start1)
	fmt.Println("end1", end1)
	fmt.Println("duration1", end1.Sub(start1).Nanoseconds())

	// make goroutine
	start2 := time.Now()
	runtime.GOMAXPROCS(2)

	ch1 := make(chan int)
	go getMaxWithGoroutine(ch1, numbers)

	ch2 := make(chan float64)
	go getAverageWithGoroutine(ch2, numbers)

	for i := 0; i < 2; i++ {
		select {
		case max := <-ch1:
			fmt.Println("hasil max", max)
		case average := <-ch2:
			fmt.Println("hasil ch2", average)
		}

	}
	end2 := time.Now()
	fmt.Println("start2", start2)
	fmt.Println("end2", end2)
	fmt.Println("duration2", end2.Sub(start2).Nanoseconds())
}

func getAverageArr(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}

func getMaxArr(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}

func getAverageWithGoroutine(ch chan float64, arr []int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- float64(sum) / float64(len(arr))
}

func getMaxWithGoroutine(ch chan int, arr []int) {
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	ch <- max
}
