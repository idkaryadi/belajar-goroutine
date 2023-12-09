package main

// wait group merupakn method dari sync yang berguna untuk "menunggu"
// sekumpulan go routine

// kalo tanpa wg, setelah semua go routine terpanggil, main() akan selesai
// untuk mengakali "Sprinln" di contoh pertama

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	// declarasi variable wg
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go doPrint(&wg, fmt.Sprintf("cetak bilangan ke::: %d", i))

		// untuk menginfokan bahwa ada penambahan go routine sebanyak 1
		wg.Add(1)
	}

	// bersifat blocking, akan menunggu hingga semua go routine yang ada di wg selesai
	wg.Wait()
}

// menggunakan alamat memori, agak "syncronise antara wg"
func doPrint(wg *sync.WaitGroup, message string) {
	// menginfokan bahwa go routine sudah berhasil dijalankan
	defer wg.Done()

	fmt.Println(message)
}
