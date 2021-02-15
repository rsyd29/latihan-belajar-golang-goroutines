package latihan_belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) { // membuat unit testing TestAtomic
	var x int64 = 0           // membuat konstanta dengan var x tipe datanya int64
	group := sync.WaitGroup{} // membuat group untuk struct WaitGroup

	for i := 1; i <= 1000; i++ { // melakukan perulangan sebanyak 1000 kali
		group.Add(1) // running 1 proses Asynchronous
		go func() {  // goroutine terhadap anonymous function
			for j := 1; j <= 100; j++ { // melakukan perulangan sebanyak 100
				//x++ // output ini mengalami race condition
				atomic.AddInt64(&x, 1) // outputnya tidak mengalami race condition sudah pasti outputnya 100000
				// tanpa harus menggunakan locking Mutex atau RWMutex
			}
			group.Done() // untuk menurunkan 1 counter
		}() // menjalankan anonymous function
	}

	group.Wait()                 // menunggu semua goroutine selesai
	fmt.Println("Counter = ", x) // cetak counter var x
}

/**
Jika terjadi error : panic: sync: WaitGroup is reused before previous Wait has returned
Itu artinya, goroutine belum selesai menjalankan kode group.Add(1), naun goroutine unit test
sudah melakukan group.Wait(), group tidak boleh di add ketika sudah di Wait(), hal ini biasanya
terjadi jika resource hardware kurang cepat ketika menjalankan goroutine diawal
Jika hal ini terjadi, silahkan pindahkan kode group.Add(1), ke baris 15 sebelum memanggil go func()
*/
