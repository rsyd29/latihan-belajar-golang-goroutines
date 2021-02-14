package latihan_belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

/**
membuat func AddToMap dengan beberapa parameter
1. data pointer dari sync.Map
2. value int
3. group pointer dari sync.WaitGroup
*/
func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()       // melakukan defer pada group agar ketika proses Add selesai maka menjalankan Add
	group.Add(1)             // melakukan penambahan 1 proses asynchronous
	data.Store(value, value) // menyimpan data ke map dengan value
}

func TestMap(t *testing.T) { // membuat unit test
	data := &sync.Map{}        // membuat sync.Map yang datanya merupakan pointer lalu dimasukkan ke dalam var data
	group := &sync.WaitGroup{} // membuat sync.WaitGroup yang datanya merupakan pointer lalu dimasukkan ke dalam var group

	for i := 0; i < 100; i++ { // melakukan perulangan goroutine sebanyak 100
		go AddToMap(data, i, group) // goroutine untuk function AddToMap dengan parameter value dengan nilai i yaitu perulangan itu sendiri
	}

	group.Wait() // menunggu proses goroutine semuanya selesai

	data.Range(func(key, value interface{}) bool { // melakukan iterasi dengan return value bernilai true
		fmt.Println(key, ":", value) // cetak key-nya
		return true                  // nilai baliknya adalah true
	})
}

/**
Jadi ini membuat Map tapi ingin diakses sama beberapa goroutine sekaligus, jadi jangan pakai map biasa dari bawaan go-lang,
melainkan menggunakan package dari bawan sync.Map karena aman dari race condition.

outputnya adalah
keynya dimulai dari 0-99 karena iterasinya sampai 100 kali, tidak terjadi race condition karena semuanya berhasil dijalankan iterasinya
41 : 41
58 : 58
67 : 67
75 : 75
--- PASS: TestMap (0.00s)
PASS
*/
