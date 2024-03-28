/* Tipe data Slice
tipe data slice adalah potongan dari tipe data array
slice mirip dengan Array, yang membedakan adalah ukuran slice bisa berubah
slice dan array selalu terkoneksi, di mana slice adalah
data yang mengakses sebagian atau seluruh data di array
simpelnya, saat membuat slice, di dalam slice ada arraynya.
hanya saja arraynya di atur oleh slice.
sehingga jika data array yang di atur oleh slice penuh,
maka slice akan membuat array baru yang otomatis di atur oleh slice

detail tipe data slice
tipe data slice memiliki 3 data, yaitu pointer, length dan capacity
pointer : penunjuk data pertama di array pada slice
length : panjang slice
capacity : kapasitas dari slice
di sini length tidak boleh melebihi dari capacity

tidak seperti array, panjang slice dapat bertambah dan
berkurang sesuai keinginan kita

beberapa cara untuk membuat slice :
1. menggunakan format slice[]{contoh}
2. membuat slice dari array
3. menggunakan function make()
*/

package main

import "fmt"

func main() {

	// contoh 1 deklarasi slice dengan length 0 dan kapasitas 0
	nomor := []int{}

	// contoh 2 deklarasi slice dengan length dan kapasitas 4
	hewan := []string{"ayam", "bebek", "kucing", "sapi"}

	// belajar slice 1

	buah := []string{"apel", "jeruk", "jambu"}

	// Contoh 3 deklarasi slice dari array
	var rak = [3]string{"buku 1", "buku 2", "buku 3"} // ini arraynya
	rakBuku1 := rak[1:3]                              // membuat slice dari array dimulai index 1 sampai index sebelum 3.
	rakBuku2 := rak[1:]                               // membuat slice dari array dimulai index 1 sampai index akhir array
	rakBuku3 := rak[:2]                               // membuat slice dari array dimulai dari index 0 sampai index sebelum 2
	rakBuku4 := rak[:]                                // membuat slice dari array dimulai dari index 0 sampai index terahir di array

	ambilSlice := rakBuku2[1]

	fmt.Println("Ini adalah Slice dengan panjang 0 dan Kapasitas 0: ", nomor)
	fmt.Println("Ini adalah Slice dengan Panjang 4 dan Kapasitas 4", hewan)
	fmt.Println("Ini adalah array yang akan dibuat Slice :", rak)
	fmt.Println(rakBuku1, "Adalah Slice yang dibuat dari Array", rak, "dimulai dari index 1 sampai index sebelum 3")
	fmt.Println(rakBuku2, "Adalah Slice yang dibuat dari Array", rak, "dimulai dari index 1 sampai akhir array")
	fmt.Println(rakBuku3, "Adalah Slice yang dibuat dari Array", rak, "dari index 0 sampai index sebelum 2")
	fmt.Println(rakBuku4, "Adalah Slice yang dibuat dari Array", rak, "dari index 1 sampai index akhir array")
	fmt.Println(ambilSlice)
	fmt.Println(len(buah))
}
