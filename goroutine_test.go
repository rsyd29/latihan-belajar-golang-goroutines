package latihan_belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	// biasanya running biasa seperti ini
	//RunHelloWorld()
	//fmt.Println("Ini langsung")
	/**
	Tanpa menggunakan goroutine keluarnya seperti ini
	=== RUN   TestCreateGoroutine
	Hello World
	Ini langsung
	--- PASS: TestCreateGoroutine (0.00s)
	PASS

	Jadi dia langsung menjalankan function RunHelloWorld() setelah itu print
	*/

	// Menjalankan goroutine
	//go RunHelloWorld() // ini akan running secara asynchronous
	//fmt.Println("Ini menggunakan goroutines")
	/**
	Menggunakan goroutines outputnya seperti ini
	=== RUN   TestCreateGoroutine
	Ini menggunakan goroutines
	--- PASS: TestCreateGoroutine (0.00s)
	PASS

	hasil function RunHelloWorld() tidak ada, kenapa? karena programnya sudah selesai sebelum
	function itu dijalankan. Jadi goroutine tersebut belum sempat untuk dieksekusi.
	*/

	// Menjalankan goroutines bersama dengan package time
	go RunHelloWorld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second) // artinya dia akan sleep selama 1 seconds / detik.
	// Untuk menunggu goroutines selesai dieksekusi

	/**
	Menggunakan goroutines bersama dengan package time, hasilnya seperti ini
	=== RUN   TestCreateGoroutine
	Ups
	Hello World
	--- PASS: TestCreateGoroutine (1.00s)
	PASS

	Jadi dia menjalankan Println terlebih dahulu, setelah itu menjalankan function
	RunHelloWorld()
	*/
}

/**
Note: Problem apabila goruotine menjalankan sebuah function yang mengembalikan sebuah nilai.
Walaupun bisa akan tetapi jadi tidak berguna, karena return valuenya tidak bisa ditangkap.
*/
