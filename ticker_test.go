package latihan_belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second) // mengambil data NewTicker selama 1 detik lalu disimpan ke dalam var ticker

	go func() { // go routine menjalankan anonymous function
		time.Sleep(5 * time.Second) // sleep selama 5 detik
		ticker.Stop()               // ticker akan stop setelah 5 detik
	}() // menjalankan anonymous function

	for time := range ticker.C { // for range untuk mendapatkan data channel dari ticker
		fmt.Println(time) // cetak time dari ticker tersebut
	}
	/**
	Outputnya
	=== RUN   TestTicker
	2021-02-16 16:24:23.941230546 +0700 WIB m=+1.000983998
	2021-02-16 16:24:24.941288152 +0700 WIB m=+2.001041623
	2021-02-16 16:24:25.941205222 +0700 WIB m=+3.000958676
	2021-02-16 16:24:26.941190034 +0700 WIB m=+4.000943423
	2021-02-16 16:24:27.941194321 +0700 WIB m=+5.000947725
	fatal error: all goroutines are asleep - deadlock!
	*/
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second) // memasukkan data ke dalam 1 channel saja

	for time := range channel { // membuat for range terhadap var channel
		fmt.Println(time) // cetak time
	}
	/**
	Outputnya
	=== RUN   TestTick
	2021-02-16 16:25:44.621215213 +0700 WIB m=+1.000563572
	2021-02-16 16:25:45.621324466 +0700 WIB m=+2.000672925
	2021-02-16 16:25:46.621178385 +0700 WIB m=+3.000526762
	2021-02-16 16:25:47.621174528 +0700 WIB m=+4.000522843
	2021-02-16 16:25:48.621195008 +0700 WIB m=+5.000543320
	2021-02-16 16:25:49.621302449 +0700 WIB m=+6.000650843

	ini tidak akan berhenti sampai kita stop timernya
	*/
}

/**
Jadi kalau ingin membuat kejadian yang berulang maka menggunakan Ticker, kalau hanya sekali saja kejadian tersebut
maka gunakan yang sebelumnya yaitu timer.
*/
