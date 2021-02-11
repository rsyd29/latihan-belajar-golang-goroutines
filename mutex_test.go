package latihan_belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Mutex
func TestMutex(t *testing.T) {
	x := 0               // variable awal dengan nilai 0 yang diberi nama x
	var mutex sync.Mutex // membuat variable mutex yang diambil dari package sync

	for i := 1; i <= 1000; i++ { // melakukan perulangan sebanyak 1000
		go func() { // membuat goroutine untuk anonymous function
			for j := 1; j <= 100; j++ { // melakukan perulangan sebanyak 100
				mutex.Lock()   // membuat mutex Lock
				x = x + 1      // increment variable x
				mutex.Unlock() // membuat mutex unlock
			}
		}() // menjalankan anonymous function
	}

	time.Sleep(5 * time.Second)  // membuat sleep sebanyak 5 detik
	fmt.Println("Counter = ", x) // mencetak counter penjumlah variable x
	/**
	Output Program
	=== RUN   TestMutex
	Counter =  100000
	--- PASS: TestMutex (5.00s)
	PASS

	Berarti ini aman, kita terhindar dari yang namanya race condition, materi sebelumnya outputnya tidak pasti.
	Sedangkan menggunakan mutex outputnya sudah pasti 100000

	Note:
	Untuk kecepatan sih lebih cepat race condition akan tetapi tidak aman,
	Sedangkan menggunakan mutex locking, maka akan lebih lama prosesnya akan tetapi lebih aman. (tapi tidak begitu lambat
	masih hitungan nanosecond), jadi lebih baik menggunakan mutex, daripada tidak sama sekali yang akan mengakibatkan
	race condition.
	*/
}
