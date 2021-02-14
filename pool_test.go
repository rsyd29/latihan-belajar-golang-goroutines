package latihan_belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{ // membuat var pool dengan nilai struct Pool
		New: func() interface{} { // New func dengan return value interface{} ini untuk membuat data baru, apabila data nil maka
			return "New" // cetak New
		},
	}

	// perintah pool.Put() untuk memasukkan datanya
	pool.Put("Budiman")
	pool.Put("Rasyid")
	pool.Put("Zainuddin")

	for i := 0; i < 10; i++ { // melakukan perulangan untuk goroutine anonymous function sebanyak 10 kali
		go func() { // goroutine anonymous function
			data := pool.Get()          // membuat var data dengan nilai diambil dari data pool
			fmt.Println(data)           // cetak data
			time.Sleep(1 * time.Second) // sleep selama 1 detik
			pool.Put(data)              // setelah tidak terpakai lagi kita masukkan kembali datanya
		}() // menjalankan anonymous function
	}

	time.Sleep(5 * time.Second) // sleep selama 5 detik
	fmt.Println("Selesai")      // cetak Selesai

	/**
	Output Program Di atas
	=== RUN   TestPool
	Budiman
	Zainuddin
	Rasyid
	New
	New
	New
	New
	New
	New
	New
	Selesai
	--- PASS: TestPool (5.00s)
	PASS
	*/
}
