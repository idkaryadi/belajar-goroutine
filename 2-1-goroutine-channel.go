package main

import "fmt"

func main() {
	// channel adalah cara berkomunikasi antar goroutine,
	// channel dideklarasikan dengan make(chan type-chane)
	// memasukkan data ke channel nama-channel <- value
	// menangkap data dari channel nama-variable <- nama-channel

	// kirim value di channel selalu dibagian yg ada go nya atau goroutine
	// kalo go routinenya ada di bagian nerima data bakal kena deadlock, karena belum selesai ngirim data udah diminta (mungkin ya)

	message := make(chan string)

	go jembatan(message, "flip")
	go jembatan(message, "evermos")
	go jembatan(message, "alterra")

	// nerima value
	// urutan nya ngacak karena bergantung pada kecepatan penyelesaian goroutine
	var nilai1 = <-message
	fmt.Println("nilai keluar jembatan::", nilai1)

	var nilai2 = <-message
	fmt.Println("nilai keluar jembatan::", nilai2)

	var nilai3 = <-message
	fmt.Println("nilai keluar jembatan::", nilai3)

}

func jembatan(c chan string, value string) {
	fmt.Println("masuk jembatan::", value)
	c <- value
}
