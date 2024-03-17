/* Function
function merupakan sebuah blok kode yang dibuat
adalam program agar bisa digunakan berulang-ulang.
cara membuat function sangat sederhana, hanya menggunakan
kata kuncu func lalu diikuti nama functionnya dan blok kode
isi functionnya
setelah membuat function, kita bisa mengeksekusi
function tersebut dengan memanggilnya menggunakan
kata kunci nama functionnya diikuti tanda kurung ().

function adalah blok pernyataan yang dapat digunakan berulang kali dalam program.
function tidak akan dijalankan secara otomatis ketika halaman dimuat.
Sebuah function akan dieksekusi dengan pemanggilan ke function tersebut.

Untuk membuat (mendeklarasikan) function, lakukan hal berikut:
Gunakan kata kunci func.
Tentukan nama untuk function tersebut, diikuti dengan tanda kurung ().
Terakhir, tambahkan kode yang mendefinisikan apa yang harus dilakukan oleh function tersebut,
di dalam kurung kurawal {}.

Syntax :
func FunctionName() {
  // kode yang akan dieksekusi
}

Aturan Penamaan function :
Nama function harus diawali dengan huruf
Nama function hanya boleh berisi karakter alfanumerik dan underscore(A-z, 0-9, dan _)
Nama function peka terhadap huruf besar/kecil
Nama fungsi tidak boleh mengandung spasi
Jika nama fungsi terdiri dari beberapa kata,
teknik yang diperkenalkan untuk penamaan variabel multi-kata dapat digunakan.

Tips: Berikan nama function yang mencerminkan apa yang dilakukan oleh function tersebut!



*/

package main

import "fmt"

// contoh function
func salam() {
	fmt.Println("Assalamu'alaikum")
}

func penjumlahan() {
	a := 3 * 4
	fmt.Println(a)
}

func benar() {
	jujur := true

	fmt.Println(jujur)
}

func ef() {

	if 100 > 200 {
		fmt.Println("100 Lebih Besar dari 200")
	} else {
		fmt.Println("100 Lebih Kecil dari 200")
	}
}

func main() {

	salam()

	penjumlahan()

	benar()

	ef()

}
