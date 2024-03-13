/* Beberapa Function untuk Slice

len(slice) : untuk mendapatkan panjang slice di mulai dari pointernya
cap(slice) : untuk mendapatkan kapasitas slice
append(slice, data) : membuat slice baru dengan menambahkan data
					  ke posisi terahir slice, jika kapasitas sudah penuh
					  maka akan membuat array baru
make([]tipedata, length, capacity) : membuat slice baru dari array yang baru
copy(destination, source) : menyalin slice dari source ke destination (dari satu slice ke slice yang lain)
*/

/* WARNING!

hati-hati saat membuat array

saat membuat array harus hati-hati, karena
jika salah bukannya membuat array tapi malah membuat slice,
atau sebaliknya

ini array ;
nomor := [...]int{1, 2, 3, 4}

ini slice ;
nomor := []int{1, 2, 3, 4}

pilih array atau slice?
ketika membuat aplikasi rata-rata menggunakan slice
jarang menggunakan array
*/

package main

import "fmt"

func main() {
	// contoh function slice

	// ini array untuk slice
	hewan := [...]string{"Ayam", "Kucing", "Burung", "Kambing"}

	// ini slice dari array
	ternak := hewan[1:] // kucing burung kambing

	// mengganti slice dari array dan arraynya juga berubah
	ternak[1] = "Kuda"
	ternak[0] = "Sapi"

	fmt.Println(hewan)
	fmt.Println(ternak)

	// membuat  slice baru
	peternakan := append(ternak, "unta")

	// mengganti data di index 0 pada slice peternakan
	peternakan[0] = "Bebek"

	fmt.Println(peternakan)
	fmt.Println(hewan) // menguji array awal

	// membuat slice baru dari dengan make

	// panjang kapasitas slice lebih banyak dari panjangnya
	// karena agar setiap kali menambahkan data ke slice
	// tidak membuat array baru, karena jika hal itu terjadi
	// aplikasi akan berjalan lambat
	// dan tidak menghemat memori
	// untuk itu, capacity dibuat lebuh banyak dari pada panjangnya
	kamar1 := make([]string, 2, 3)
	kamar1[0] = "zaid"
	kamar1[1] = "umar"
	// kamar1[2] = "bakar" ini error karena untuk menambahkan dengan append, soalnya
	// di atas sudah di tentuntakan panjangnya 2

	kamar2 := append(kamar1, "bakar")

	fmt.Println(kamar1)
	fmt.Println(len(kamar1))
	fmt.Println(cap(kamar1))
	fmt.Println(kamar2)

	// kode program copy slice

	sawahan := []string{"Rt 07", "Rt 08", "Rt 09"}
	pundung := make([]string, len(sawahan), cap(sawahan))

	copy(pundung, sawahan)

	fmt.Println(pundung)
}
