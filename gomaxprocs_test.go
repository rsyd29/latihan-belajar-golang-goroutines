package latihan_belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{} // membuat struct WaitGroup

	for i := 0; i < 100; i++ { // melakukan perulangan sebanyak 100 kali untuk menjalankan goroutine
		group.Add(1) // melakukan proses penambahan 1 asynchronous
		go func() {  // membuat goroutine terhadap anonymous function
			time.Sleep(3 * time.Second) // melakukan sleep 3 detik
			group.Done()                // melakukan proses pengurangan dari goroutine
		}()
	}

	totalCpu := runtime.NumCPU()        // untuk mengetahui jumlah CPU komputer yang digunakan
	fmt.Println("Total CPU:", totalCpu) // mencetak totalCpu

	totalThread := runtime.GOMAXPROCS(-1)     // untuk mengetahui jumlah Thread komputer yang digunakan
	fmt.Println("Total Thread:", totalThread) // mencetak totalThread

	totalGoroutine := runtime.NumGoroutine()        // untuk mengetahui jumlah goroutine yang sedang berjalan
	fmt.Println("Total Goroutine:", totalGoroutine) // mencetak totalGoroutine

	group.Wait() // menunggu semua goroutine selesai

	/**
	Outputnya sebelum membuat goroutine
	=== RUN   TestGetGomaxprocs
	Total CPU: 4
	Total Thread: 4
	Total Goroutine: 2
	--- PASS: TestGetGomaxprocs (0.00s)
	PASS
	*/

	/**
	Outputnya setelah membuat goroutine
	=== RUN   TestGetGomaxprocs
	Total CPU: 4
	Total Thread: 4
	Total Goroutine: 102
	--- PASS: TestGetGomaxprocs (3.00s)
	PASS
	*/
}

// Mengubah Jumlah Thread
func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCpu)

	runtime.GOMAXPROCS(20)                // untuk mengubah jumlah thread, ini jarang sekali dilakukan perubahan thread, karena golang sendiri sudah optimal sekali dalam melakukan manajemen thread
	totalThread := runtime.GOMAXPROCS(-1) // parameternya kalau > 0 maka akan mengubah jumlah thread-nya, sedangkan < 0 itu tidak mengubah thread-nya
	fmt.Println("Total Thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoroutine)

	group.Wait()
	/**
	Outputnya
	=== RUN   TestChangeThreadNumber
	Total CPU: 4
	Total Thread: 20
	Total Goroutine: 102
	--- PASS: TestChangeThreadNumber (3.00s)
	PASS
	*/
}

/**
Jadi tergantung kebutuhan kita, dalam membutuhkan Thread-nya, jadi lebih baik tidak usah mengubah thread-nya Go-Lang
karena defaultnya sudah optimal, kalau mau cepat lagi prosesnya kita bisa menggunakan horizontal scalling dibanding
mengubah thread-nya secara vertikal.

Intinya tidak perlu manage Thread dan Goroutine-nya secara manual, itu semuanya sudah diatur oleh go scheduler.
*/
