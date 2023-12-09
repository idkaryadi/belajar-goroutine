package main

import (
	"fmt"
	"runtime"
	"sync"
)

// race condition: kondisi ketika 1 data diakses bareng bareng sama banyak
// go routine, jadinya yg paling cepat yg bisa akses, yg lain gak bisa.
// hasilnya jadi kacau

// cara ngecek apakah ada race condition di go
// go run -race main.go

type counter struct {
	val int
}

func (c *counter) Add() {
	c.val++
}

func (c *counter) Value() int {
	return c.val
}

func main() {
	// harus >1 agar terjadi race condition
	// makin besar makin banyak race condition
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				meter.Add()
			}
		}()
		wg.Add(1)
	}

	wg.Wait()
	// kalo tidak ada race condition harusnya jadi 1000
	fmt.Println("result::", meter.Value())

}
