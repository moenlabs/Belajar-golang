/* Recover
recover merupakan sebuah function yang bisa digunakan untuk menangkap panic
ketika terjadi error, kita bisa mengecek errornya kenapa dengan function recover
dengan recover juga, proses panic akan terhenti, sehingga program tetap akan berjalan
*/

package main

import "fmt"

// contoh 1 cara salah menggunakan recover
/*
func berakhir() {
	fmt.Println("Aplikasi Sudah Berakhir")
}
func aplikasi(error bool) {
	defer berakhir()

	if error {
		panic("Terjadi Error")
	} // kode setelah scope if tidak akan di jalankan

	pesan := recover()
	fmt.Println("Terjadi Panic dan direcover", pesan)
}

*/

// contoh yang benar dengan menyimpan function recover() di function defer

func penggunaanRecover() {
	fmt.Println("Aplikasi Sudah Berakhir")
	pesan := recover()
	fmt.Println("Terjadi Error dan direcover", pesan)
}

func percobaanError(error bool) {
	defer penggunaanRecover()
	if error {
		panic("\nTerjadi Panic")
	}
}
func main() {
	//aplikasi(true)
	percobaanError(true)
}
