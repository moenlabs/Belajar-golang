/* Pointer di Golang
Saat kita membuat parameter di function, secara default
adalah pass by value, artinya data akan di copy lalu di kirim ke function tersebut.
oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya
tidak akan pernah berubah
hal ini membuat variable menjadi aman, karena tidak bisa diubah
namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
untuk melakukan ini, kita juga bisa menggunakan pointer di function.
untuk menjadikan sebuah parameter sebagai pointer,
kita bisa menggunakan operator * di parameternya.
*/

package main

import "fmt"

type Alamat struct {
	Prov, Kab, Kec string
}

// contoh parameter tanpa pointer
func GantiAlamatJogja(alamat Alamat) {
	alamat.Prov = "Jogja"
}

// contoh pakai pointer
func AlamatBaru(provence *Alamat) {
	provence.Prov = "Yogyakarta"
}

func main() {
	alamat := Alamat{"Jateng", "Magelang", "Salam"}

	GantiAlamatJogja(alamat) // memanggil function yang tidak pakai pointer

	rumah := Alamat{"Magelang", "Sleman", "Gamping"}

	AlamatBaru(&rumah) // harus dikasih tanda &

	fmt.Println(alamat) // profinsi ridak berubah
	fmt.Println(rumah)  // profinsi yang tadinya kosong di ganti Yogyakarta
}
