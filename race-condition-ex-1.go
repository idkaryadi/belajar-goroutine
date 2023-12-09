package main

import "fmt"

func main() {
	done := make(chan bool)
	m := make(map[string]string)
	m["name"] = "world"
	go func() {
		m["name"] = "data race"
		done <- true
	}()
	// kena race condition karena m["name"] diubah di go routine
	fmt.Println("Hello,", m["name"])
	<-done
}
