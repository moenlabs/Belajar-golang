/*
Pass by Value
Secara default di golang semua variable itu
passing by value, bukan pass by reference.
passing by value artinya, jika kita mengirim sebuah variabel
ke dalam function, method atau variabel lain
sebenarnya yang kita kirim adalah duplikasinya.
sehingga data di rubah, data awal akan aman,
karena data yang ke dua diduplikasi

Pointer
jika ingin membuat data pass by reference maka itu akan mengunakan pointer.
Pointer adalah kemampuan membuat reference ke lokasi data di memoty yang sama,
tanpa menduplikasi data yang sudah ada.
sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

Operator &
untuk membuat variabel dengan nilai pointer
ke variable yang lain, kita bisa menggnakan operator &
diikuti nama variabelnya
*/

package main

import "fmt"

// struct untuk contoh pass by value (tidak menggunakan pointer)
type Alamat struct {
	Prov, Kab, Kec string
}

// struct contoh pass by reference (menggunakan pointer)
type Address struct {
	City, Provence, State string
}

func main() {
	alamat1 := Alamat{"Jogja", "Sleman", "Gamping"}
	alamat2 := alamat1

	alamat2.Kec = "Mlati"

	fmt.Println(alamat1) // alamat 1 tidak berubah
	fmt.Println(alamat2) // alamat 2 berubah karena data alamat 1 di copy atau disebut dengan pass by value

	address1 := Address{"Sleman", "Yogyakarta", "Indonesia"}
	address2 := &address1 // menggunakan operator & untuk merujuk ke Address

	address2.City = "Bantul"

	fmt.Println(address1)
	fmt.Println(address2)

}
