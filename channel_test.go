package latihan_belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // membuat channel dengan tipe data string
	// setelah membuat channel kita close dengan menggunakan defer, jadi defer ini mau error atau tidak tetap di close
	defer close(channel)

	/*
		// untuk mengirim data
		channel <- "Budiman Rasyid"

		// untuk menerima data dari channel
		data := <- channel
		fmt.Println(data)

		// kalau tidak mau disimpan ke dalam variable bisa langsung seperti ini
		fmt.Println(<- channel)
	*/

	go func() { // membuat goroutine dan anonymous function
		time.Sleep(2 * time.Second)            // sleep selama 2 detik
		channel <- "Budiman Rasyid Zainuddin " // lalu kirim datanya ke channel
		fmt.Println("Selesai Mengirim Data ke Channel")
	}() // untuk mengeksekusinya dengan menggunakan ()

	/**
	Jadi apabila data ada yang blm diterima maka channel di atas akan nge-block (diam)
	sampai channel tersebut yang di atas telah diambil. Apabila tidak diambil maka dia akan
	diam saja sampai programnya stop.

	lalu apabila di anonymous function tersebut tidak ada channel, lalu menerima data dari
	channel yang sebenarnya tidak ada maka terjadi fatal error, karena deadlock artinya
	tidak ada channel yang ada.

	Jadi apabila ingin mengirim data pastikan ada yang menerima datanya, lalu apabila
	ingin menerima pastikan ada data yang akan dikirim oleh channel tersebut.
	*/

	data := <-channel // ini untuk menunggu datanya dan dimasukkan ke dalam var data
	fmt.Println(data) // setelah itu kita lihat datanya apa, apakah sama yang diterima sama yang dikirmnya atau tidak

	time.Sleep(5 * time.Second)

	/**
	Output Program
	=== RUN   TestCreateChannel
	Budiman Rasyid Zainuddin
	Selesai Mengirim Data ke Channel
	--- PASS: TestCreateChannel (7.00s)
	PASS
	*/
}
