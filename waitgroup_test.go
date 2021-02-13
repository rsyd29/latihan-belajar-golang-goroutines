package latihan_belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // proses defer berguna untuk apabila ada proses yang gagal maka method Done ini akan tetap dijalankan

	group.Add(1) // running 1 proses Asynchronous

	// melakukan proses
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait() // ini berguna untuk menunggu semua goroutine dijalankan
	fmt.Println("Selesai")
}

/**
Note: Ketika membuat WaitGroup pastikan proses Wait tersebut harus di Done, agar tidak terjadi deadlock.
Jadi ketika ingin menunggu proses goroutine selesai semua, daripada menggunakan sleep lebih baik menggunakan WaitGroup.

outputnya
...
Hello
Hello
Hello
Hello
Selesai
--- PASS: TestWaitGroup (1.00s)
PASS
*/
