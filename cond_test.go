package latihan_belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}        // variable locker dengan nilai struct Mutex (untuk melakukan locking)
var cond = sync.NewCond(&locker) // variable cond untuk membuat struct NewCond dengan parameter pointer locker
var group = sync.WaitGroup{}     // variable group untuk membuat WaitGroup (untuk menunggu semua proses goroutine selesai)

func WaitCondition(value int) { // function WaitCondition dengan parameter value
	defer group.Done() // proses defer berguna untuk apabila ada proses yang gagal maka method Done ini akan tetap dijalankan
	group.Add(1)       // running 1 proses Asynchronous

	cond.L.Lock()              // melakukan Locking terhadap cond (condition)
	cond.Wait()                // untuk menunggu apakah perlu menunggu atau tidak dari goroutine yang ada
	fmt.Println("Done", value) // mencetak
	cond.L.Unlock()            // melakukan Unlocking terhadap cond (condition)
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i) // menjalankan function WaitCondition dengan goroutine sebanyak 10 kali
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // menjalankan satu-satu goroutine
		}
	}()
	/**
	Output Program Menggunakan Signal(), dia jalan satu-satu goroutine-nya dengan concurrent & Parallel oleh karena itu
	tidak berurutan (random)
	=== RUN   TestCond
	Done 1
	Done 0
	Done 4
	Done 2
	Done 3
	Done 6
	Done 5
	Done 7
	Done 8
	Done 9
	--- PASS: TestCond (10.00s)
	PASS
	*/

	//go func() { // menjalankan anonymous function menggunakan goroutine
	//	time.Sleep(1 * time.Second) // sleep selama 1 detik
	//	cond.Broadcast() // menjalankan semua goroutine
	//}()
	/**
	Output Program Menggunakan Broadcast, dia akan menjalankan semua goroutine dalam 1 detik sekaligus,
	=== RUN   TestCond
	Done 9
	Done 3
	Done 4
	Done 7
	Done 6
	Done 8
	Done 5
	Done 1
	Done 2
	Done 0
	--- PASS: TestCond (1.00s)
	PASS
	*/

	group.Wait() // ini berguna untuk menunggu semua goroutine dijalankan
}

/**
Note:
Ketika ingin membutuhkan locking tapi ingin lock-nya jalan, ketika kalian perintahkan jalan maka bisa menggunakan
sync.Cond ini
*/
