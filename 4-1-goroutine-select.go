package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	numbers := []int{1, 21, 43, 33, 65, 19, 22, 54, 31, 28, 37, 3, 4, 8, 11}

	ch1 := make(chan int)
	go getMax(ch1, numbers)

	ch2 := make(chan float64)
	go getAverage(ch2, numbers)

	// before select, uncomment if need
	// average := <-ch2
	// fmt.Println("print average", average)

	// max := <-ch1
	// fmt.Println("print max", max)

	// after select
	// kalo gak ada loopingnya, program bakal selesai ketika salah satu goroutine selesai
	for i := 0; i < 2; i++ {
		select {
		case max := <-ch1:
			fmt.Println("hasil max", max)
		case average := <-ch2:
			fmt.Println("hasil ch2", average)
		}

	}
}

func getAverage(ch chan float64, arr []int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- float64(sum) / float64(len(arr))
}

func getMax(ch chan int, arr []int) {
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	ch <- max
}
