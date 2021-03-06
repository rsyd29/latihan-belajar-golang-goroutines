package latihan_belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // membuat channel dengan tipe data string
	// setelah membuat channel kita close dengan menggunakan defer, jadi defer ini mau error atau tidak tetap di close
	defer close(channel)

	/*
		// untuk mengirim data
		channel <- "Budiman Rasyid"

		// untuk menerima data dari channel
		data := <- channel
		fmt.Println(data)

		// kalau tidak mau disimpan ke dalam variable bisa langsung seperti ini
		fmt.Println(<- channel)
	*/

	go func() { // membuat goroutine dan anonymous function
		time.Sleep(2 * time.Second)            // sleep selama 2 detik
		channel <- "Budiman Rasyid Zainuddin " // lalu kirim datanya ke channel
		fmt.Println("Selesai Mengirim Data ke Channel")
	}() // untuk mengeksekusinya dengan menggunakan ()

	/**
	Jadi apabila data ada yang blm diterima maka channel di atas akan nge-block (diam)
	sampai channel tersebut yang di atas telah diambil. Apabila tidak diambil maka dia akan
	diam saja sampai programnya stop.

	lalu apabila di anonymous function tersebut tidak ada channel, lalu menerima data dari
	channel yang sebenarnya tidak ada maka terjadi fatal error, karena deadlock artinya
	tidak ada channel yang ada.

	Jadi apabila ingin mengirim data pastikan ada yang menerima datanya, lalu apabila
	ingin menerima pastikan ada data yang akan dikirim oleh channel tersebut.
	*/

	data := <-channel // ini untuk menunggu datanya dan dimasukkan ke dalam var data
	fmt.Println(data) // setelah itu kita lihat datanya apa, apakah sama yang diterima sama yang dikirmnya atau tidak

	time.Sleep(5 * time.Second)

	/**
	Output Program
	=== RUN   TestCreateChannel
	Budiman Rasyid Zainuddin
	Selesai Mengirim Data ke Channel
	--- PASS: TestCreateChannel (7.00s)
	PASS
	*/
}

// Channel Sebagai Parameter video 91
func GiveMeResponse(channel chan string) { // membuat function dengan parameter channel
	// jadi parameter di atas tidak perlu lagi pointer, karena langsung reference data aslinya
	time.Sleep(2 * time.Second) // artinya akan sleep selama 2 detik
	channel <- "Budiman Rasyid" // ini mengirim data ke dalam channel
	//fmt.Println("Selesai Mengirim Data ke Channel")
}

func TestChannelAsParameter(t *testing.T) { // membuat func unit test sebagai parameter
	channel := make(chan string) // membuat channel lalu disimpan ke dalam var channel
	// setelah membuat channel kita close dengan menggunakan defer, jadi
	// defer ini mau error atau tidak tetap di close
	defer close(channel)

	go GiveMeResponse(channel) // menjalankan goroutine untuk function dengan parameter
	data := <-channel          // ini untuk menerima data dari channel ke dalam var data
	fmt.Println(data)          // untuk menampilkan isi data tersebut
	fmt.Println("Channel berhasil mengularkan datanya")

	time.Sleep(5 * time.Second) // artinya akan sleep selama 5 detik, untuk menunggu goroutine diatas
}

// Channel In dan Out video 92
/**
ini hanya boleh mengirim channel dengan menambahkan <- setelah chan
contoh func OnlyIn
*/
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)           // artinya akan sleep selama 2 detik
	channel <- "Budiman Rasyid Zainuddin" // ini mengirim data ke dalam channel
	// data := <- channel // error: receive from send-only type chan<- string
}

/**
ini hanya boleh menerima channel dengan menambahkan <- sebelum chan
contoh func OnlyOut
*/
func OnlyOut(channel <-chan string) {
	// channel <- "Budiman Rasyid Zainuddin" // error: channel <- "Budiman Rasyid Zainuddin" (send to receive-only type <-chan string
	data := <-channel // ini untuk menerima data dari channel ke dalam var data
	fmt.Println(data) // untuk menampilkan isi data tersebut
}

// membuat testing
func TestInOutChannel(t *testing.T) {
	channel := make(chan string) // membuat channel dengan tipe data string, lalu dimasukkan ke dalam var channel
	defer close(channel)         // untuk close channel

	go OnlyIn(channel)  // menjalankan goroutine untuk function OnlyIn dengan parameter channel
	go OnlyOut(channel) // menjalankan goroutine untuk function OnlyOut dengan parameter channel

	time.Sleep(5 * time.Second) // untuk sleep selama 5 detik
}

/**
Note:
Jadi apabila kita paksa apabila kita memasukkan ke dalam function OnlyIn
dimana function tersebut memiliki parameter in, lalu kita masukkan
data := <- channel yang artinya itu untuk menerima data yang seharusnya hanya bisa
pada function OnlyOut maka akan terjadi Error, dan sebaliknya.
*/

// Buffered Channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() { // membuat goroutine dengan anonymous function untuk mengirim data ke channel
		// jadi apabila kita blm membuat buffer di channel tersebut maka
		// akan terjadi error deadlock
		channel <- "Budiman"
		channel <- "Rasyid"
		channel <- "Zainuddin"
	}() // () untuk menjalankan anonymous function

	go func() { // membuat goroutine dengan anonymous function untuk menerima data dari channel
		// selanjutnya kita ambil channel tersebut
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		// lalu apabila datanya diambil lagi, padahal datanya sudah kosong, maka akan terjadi error
		// fmt.Println(<- channel)
	}() // () untuk menjalankan anonymous function

	time.Sleep(2 * time.Second) // sleep selama 2 detik untuk menunggu goroutine dijalankan
	fmt.Println("Selesai")
	// lalu apabila kita menambahkan buffer, maka seakan-akan datanya itu
	// masuk ke dalam slot buffer terlebih dahulu
}

/**
NOTE:
Jadi ada sedikit berbeda sama channel yang biasa, kalau channel biasa
apabila datanya masuk ke dalam channel, itu karena dia tidak punya
buffer, maka otomatis dia diminta untuk menunggu sampai ada yang
mengambil.

Kalau channel tersebut ditambahkan buffer itu otomatis akan masuk
ke dalam slot channel buffer, jadi tidak perlu menunggu lagi, kecuali
slot buffernya sudah penuh tidak ada yang kosong, baru diminta untuk
menunggu.
*/

// Range Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string) // membuat channel tanpa buffered channel
	/**
	tanpa range channel
	data := <- channel
	data := <- channel
	dst
	pertanyaanya mau berapa kali akses memasukkan data channel ke dalam var data?
	*/
	// membuat goroutine untuk anonymous function
	go func() {
		// membuat perulangan sampai 10 kali
		for i := 0; i < 10; i++ {
			// akan mengirim datanya ke dalam channel
			channel <- "Perulangan ke " + strconv.Itoa(i) // strconv untuk konversi dari int ke string
		}
		// setelah mengirim kita akan close channel
		close(channel)
		// kalau tidak di close maka perulangan data yang di bawah maka tidak akan pernah berhenti
	}()
	// daripada seperti di atas lebih baik seperti ini lakukan perulangan
	for data := range channel {
		fmt.Println("Menerima data ", data)
	}
	// kita tidak perlu sleep karena data diatas akan diulang terus sampai close
	fmt.Println("Selesai")
}

// Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string) // membuat channel lalu diberi nama ke var channel1
	channel2 := make(chan string) // membuat channel lalu diberi nama ke var channel2
	defer close(channel1)         // setelah dikirim akan close channel1
	defer close(channel2)         // setelah dikirim akan close channel2

	go GiveMeResponse(channel1) // Membuat goroutines dan masukkan ke dalam func GiveMeResponse dengan parameter channel1
	go GiveMeResponse(channel2) // Membuat goroutines dan masukkan ke dalam func GiveMeResponse dengan parameter channel2

	/**
	Jadi ini mengambil data yang paling cepat datang terlebih dahulu
	maka dia yang akan diambil datanya
	*/
	// Select ke-1
	//select { // melakukan select channel
	//case data := <-channel1: // mengambil data dari channel1 lalu disimpan ke dalam var data
	//	fmt.Println("Data dari Channel 1", data) // setelah itu di print datanya
	//case data := <-channel2: // mengambil data dari channel2 lalu disimpan ke dalam var data
	//	fmt.Println("Data dari Channel 2", data) // setelah itu di print datanya
	//}
	/**
	Output
	=== RUN   TestSelectChannel
	Selesai Mengirim Data ke Channel
	Data dari Channel 1 Budiman Rasyid
	--- PASS: TestSelectChannel (2.00s)
	PASS

	Channel 2 nya tidak ke ambil, karena kita hanya melakukan select sekali
	*/

	// Select ke-2
	//select { // melakukan select channel
	//case data := <-channel1: // mengambil data dari channel1 lalu disimpan ke dalam var data
	//	fmt.Println("Data dari Channel 1", data) // setelah itu di print datanya
	//case data := <-channel2: // mengambil data dari channel2 lalu disimpan ke dalam var data
	//	fmt.Println("Data dari Channel 2", data) // setelah itu di print datanya
	//}

	// kalau channelnya banyak lebih baik menggunakan perulangan, seperti kodingan di bawah ini
	// kalau tanpa counter dia akan melakukan infinite loop, looping yang berulang-ulang. Maka akan terjadi deadlock
	// oleh karena itu buatlah counter
	counter := 0
	for { // perulangan
		select {
		case data := <-channel1: // mengambil data dari channel1 lalu disimpan ke dalam var data
			fmt.Println("Data dari Channel 1", data)
			counter++ // Increment Counter
		case data := <-channel2: // mengambil data dari channel2 lalu disimpan ke dalam var data
			fmt.Println("Data dari Channel 2", data)
			counter++ // Increment Counter
		}

		if counter == 2 { // apabila Counter sudah sampai 2 maka
			break // perulangan akan berhenti
		}
	}
	/**
	Output Program Select Channel menggunakan Counter Perulangan
	=== RUN   TestSelectChannel
	Data dari Channel 1 Budiman Rasyid
	Data dari Channel 2 Budiman Rasyid
	--- PASS: TestSelectChannel (2.00s)
	PASS
	*/
}

// Default Select
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string) // membuat channel lalu diberi nama ke var channel1
	channel2 := make(chan string) // membuat channel lalu diberi nama ke var channel2
	defer close(channel1)         // setelah dikirim akan close channel1
	defer close(channel2)         // setelah dikirim akan close channel2

	go GiveMeResponse(channel1) // Membuat goroutines dan masukkan ke dalam func GiveMeResponse dengan parameter channel1
	go GiveMeResponse(channel2) // Membuat goroutines dan masukkan ke dalam func GiveMeResponse dengan parameter channel2

	counter := 0
	for { // perulangan
		select {
		case data := <-channel1: // mengambil data dari channel1 lalu disimpan ke dalam var data
			fmt.Println("Data dari Channel 1", data)
			counter++ // Increment Counter
		case data := <-channel2: // mengambil data dari channel2 lalu disimpan ke dalam var data
			fmt.Println("Data dari Channel 2", data)
			counter++ // Increment Counter
		default: // default apabila belum ada datanya
			fmt.Println("Menunggu Data") // maka akan menampilkan ini
		}

		if counter == 2 { // apabila Counter sudah sampai 2 maka
			break // perulangan akan berhenti
		}
	}
}

/**
Note:
Jadi default ini untuk menunggu data apabila datanya belum masuk ke dalam select channel, maka akan dijalankan
sebuah default, sampai datanya sudah ada di select channel.

Output Program Default Select
=== RUN   TestDefaultSelectChannel
Menunggu Data
Menunggu Data
sampai 13K untuk menunggu data
Data dari Channel 1 Budiman Rasyid
Data dari Channel 2 Budiman Rasyid
--- PASS: TestDefaultSelectChannel (2.16s)
PASS
*/
