package latihan_belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)     // Membuat timer 5 detik waktu lamanya
	fmt.Println("Waktu Sekarang =", time.Now()) // lalu kita cetak waktu sekarang

	time := <-timer.C                               // setelah itu kita mengambil data dari channel timer.C (channel-nya timer) lalu simpan ke dalam var timer.
	fmt.Println("Waktu sekarang + 5 detik =", time) // lalu data time tersebut dicetak

	/**
	fmt.Println di atas berfungsi untuk membandingkan waktu sekarang dengan waktu yang ditambahkan oleh NewTimer yaitu 5
	detik
	Outputnya
	=== RUN   TestTimer
	Waktu Sekarang =
	 2021-02-16 09:41:27.450495612 +0700 WIB m=+0.000584132
	Waktu sekarang + 5 detik =
	 2021-02-16 09:41:32.450662063 +0700 WIB m=+5.000750685
	--- PASS: TestTimer (5.00s)
	PASS
	*/
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second) // ini langsung masukkan ke dalam channel
	fmt.Println(time.Now())                // cetak waktu sekarang

	time := <-channel // mengambil data dari channel lalu dimasukkan ke dalam var time
	fmt.Println(time) // cetak data time tersebut
	/**
	Outputnya sama seperti TestTimer, yang membedakan ini langsung
	=== RUN   TestAfter
	2021-02-16 09:48:22.883495201 +0700 WIB m=+0.001153522
	2021-02-16 09:48:27.883582485 +0700 WIB m=+5.001240825
	--- PASS: TestAfter (5.00s)
	PASS
	*/
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{} // Memasukkan struct WaitGroup ke dalam var group
	group.Add(1)              // untuk menambahkan 1 proses asynchronous

	time.AfterFunc(5*time.Second, func() { // membuat timer selama 5 detik, setelah itu menjalankan sebuah function
		// isi function cetak waktu sekarang + 5 detik
		fmt.Println("Execute after 5 second: ",
			time.Now())
		group.Done() // untuk mengakhir 1 proses asynchronous
	})
	fmt.Println("Execute Now: ", time.Now()) // cetak waktu sekarang
	group.Wait()                             // menunggu semua proses
	/**
	Outputnya adalah
	=== RUN   TestAfterFunc
	Execute Now:  2021-02-16 10:01:22.040242188 +0700 WIB m=+0.000528305
	Execute after 5 second:  2021-02-16 10:01:27.040446851 +0700 WIB m=+5.000733146
	--- PASS: TestAfterFunc (5.00s)
	PASS
	*/
}
