/* Tipe data map
map adalah tipe data lain yang berisikan kumpulan data yang sama,
namun bisa ditentukan jenis tipe data index yang akan digunakan
atau map adalah tipe data kumpulan key-value (kata kunci dan linalinya)
di mana kata kuncinya harus unik dan tidak boleh sama
adapun jumlah data yang di masukkan dalam map boleh sebanyak-banyaknya, asalkan kata kuncinya berbeda
jika kata kunci digunakan untuk nilai yang sama
maka secara otomatis data sebelumnya diganti dengan data yang baru.
Map adalah koleksi yang tidak berurutan dan dapat diubah
sehingga tidak memungkinkan adanya duplikasi.
Panjang map adalah jumlah elemennya.
untuk mengetahui panjangnya dengan menggunakan fungsi len().
Nilai default map adalah nil.
map menyimpan referensi ke tabel hash yang mendasarinya.
*/

package main

import "fmt"

func main() {

	//contoh 1 deklarasi map menggunakan var
	var orang = map[string]string{"nama": "zaid", "alamat": "Sawahan", "Hobi": "Ngoding"}

	// contoh 2 deklarasi map menggunakan var
	var siswa = map[string]string{
		"name 1": "Andi",
		"name 2": "Tono",
		"name 3": "Fajar", // harus pakai koma di akhirnya
	}
	// deklarasi map menggunakan :=
	data := map[string]int{"Anggur": 10, "Apel": 5, "Jeruk": 7}

	// contoh map kosong
	var kamar = map[string]int{}

	// menambahkan key dan value ke map
	kamar["Jumlah Anak"] = 100
	kamar["jumlah uang"] = 10000

	// belajar map lagi

	barang := map[string]int{"Sendok": 15, "Piring": 10, "Gelas": 5}

	fmt.Println(orang)
	fmt.Println(siswa)
	fmt.Println(data)
	fmt.Println(kamar)
	fmt.Println(barang)
	fmt.Println(barang["Piring"])

	// jika mengangakses key yang tidak ada
	// maka hasilnya adalah default dari value dari key tersebut

	fmt.Println(kamar["Jumlah Sandal"]) // hasil 0 karena yang diprint adalah default dari valuenya
}
