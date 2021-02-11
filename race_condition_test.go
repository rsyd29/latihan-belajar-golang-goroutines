package latihan_belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Race Condition
func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
				//fmt.Printf("x = %d + 1\n", x)
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter : ", x)
	/**
	Outputnya berubah rubah x nya, kenapa bisa seperti itu? kita punya satu variable yaitu variable x, dimana variable
	itu di sharing ke beberapa goroutine bahkan sampai 1000 goroutine yang mengakses variable yang sama yaitu untuk
	melakukan increment pada variable x. Lalu dimanipulasi variable tersebut yaitu var x. Jadi bisa saja ada 1 goroutine
	mengakses data dengan value x nya yang akhirannya itu tetap sama juga.
	Misal ada goroutine mengakses
	x = 1000 + 1
	lalu ada goroutine lain mengakses itu pula yaitu x = 1000 + 1 juga
	karena ini kan parallel, lalu bagaimana ada 15K data yang sama untuk menaikkan counter secara bersamaan.
	Maka itu sangat berbahaya sekali.
	Oleh karena itu solusinya adalah di materi selanjutnya
	*/
}
