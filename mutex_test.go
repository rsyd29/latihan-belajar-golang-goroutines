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

// RWMutex
/**
Kasusnya ketika membuat struct lalu variable tersebut akan diakses sekaligus oleh go routine, jadi lebih baik menggunakan
RWMutex
*/
type BankAccount struct { // membuat struct dengan nama BankAccount
	// yang memiliki field diantaranya
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) { // function untuk menambahkan balance
	account.RWMutex.Lock()                     // untuk proses menulis Lock
	account.Balance = account.Balance + amount // penjumlahan untuk balance
	account.RWMutex.Unlock()                   // untuk proses menulis unlock
}

func (account *BankAccount) GetBalance() int { // function untuk mengambil balance dengan return value int
	account.RWMutex.RLock()    // untuk proses membaca lock
	balance := account.Balance // untuk mengambil balance dan disimpan ke var balance
	account.RWMutex.RUnlock()  // untuk proses membaca unlock
	return balance             // mengembalikan nilai berupa balance dengan tipe data int
}

func TestRWMutex(t *testing.T) { // membuat unit test dengan nama TestRWMutex
	account := BankAccount{} // membuat variable account dengan nilai struct dari BankAccount

	for i := 0; i < 100; i++ { // melakukan perulangan untuk goroutine anonymous function sebanyak 100
		go func() { // goroutine untuk anonymous function
			for j := 0; j < 100; j++ { // perulangan untuk menambahkan balance sebanyak 100 kali perulangan
				account.AddBalance(1)             // menambahkan balance sebesar 1 setiap perulangan melalui method AddBalance()
				fmt.Println(account.GetBalance()) // mencetak balance yang telah diambil melalui method GetBalance()
			}
		}() // menjalnak anonymous function
	}

	time.Sleep(5 * time.Second)                           // sleep selama 5 detik
	fmt.Println("Total Balance : ", account.GetBalance()) // akan mencetak total balance-nya
	/**
	Outputnya berurutan tidak terjadi race condition
	9997
	9998
	9999
	10000
	Total Balance :  10000
	--- PASS: TestRWMutex (5.00s)
	PASS

	Kalau misalnya RWMutex diberi komentar, dimana ada data yang mengalami race condition. Hasilnya tidak berurutan
	dan terkadang total balance bisa mencapai 10000 dan bisa kurang
	9995
	9996
	9997
	9998
	Total Balance :  9998
	--- PASS: TestRWMutex (5.00s)
	PASS

	Process finished with exit code 0
	*/
}
