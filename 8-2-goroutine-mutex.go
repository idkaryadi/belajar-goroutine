package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type counter2 struct {
	sync.Mutex
	val int
}

// best practice nya
func (c *counter2) Add() {
	// lock (dari mutex) aga eksklusif hanya bisa diakses 1 goroutine
	c.Lock()
	c.val++
	// setelah selesai diunlock
	c.Unlock()
}

func (c *counter2) Value() int {
	return c.val
}

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter2

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			meter.Add()
		}()
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println("result::", meter.Value())
	end := time.Now()

	fmt.Println("DURATUON", end.Sub(start).Nanoseconds())
}
