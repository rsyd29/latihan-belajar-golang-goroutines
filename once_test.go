package latihan_belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0 // initial var counter sebagai konstanta dengan nilai 0

func OnlyOnce() { // membuat func OnlyOnce tanpa parameter dan return value
	counter++
}

func TestOnce(t *testing.T) { // membuat unit testing
	once := sync.Once{}       // memanggil Once dengan var once
	group := sync.WaitGroup{} // memanggil WaitGroup dengan var group

	for i := 0; i < 100; i++ { // membuat perulangan sebanyak 100
		go func() { // membuat anonymous function dengan goroutine
			group.Add(1)      // running 1 proses asynchronous
			once.Do(OnlyOnce) // ini bukan memanggil function melainkan nama functionnya, dan ingat hanya bisa function tanpa parameter
			/**
			Proses di atas outputnya menggunakan Once, jadi function yang sama akan dijalankan satu kali. Karena goroutine yang lain dihiraukan.
			apabila hanya boleh 1 kali saja goroutine dijalankan.
			=== RUN   TestOnce
			Counter 1
			--- PASS: TestOnce (0.00s)
			PASS
			*/

			//OnlyOnce()
			/**
			Proses di atas outputnya berubah ubah kalau tanpa Once, mengalami Race Condition
			=== RUN   TestOnce
			Counter 91
			--- PASS: TestOnce (0.00s)
			PASS
			*/
			group.Done() // menyelesaikan proses asynchronous
		}() // menjalankan anonymous function
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
