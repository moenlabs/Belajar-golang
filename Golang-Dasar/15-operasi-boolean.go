/* Operasi Bolean adalah operasi untuk menentukan logika antar variabel atau nilai
yaitu :

&& (DAN) : Mengembalikan nilai true jika kedua pernyataan tersebut benar
|| (ATAU) : Mengembalikan nilai true jika salah satu pernyataan benar
! (Tidak) : Membalikkan hasil, mengembalikan nilai salah jika hasilnya benar
*/

// Contoh kode

package main

import "fmt"

func main() {

	nilaiAkhir := 80
	nilaiAbsen := 75

	lulusNilaiAkhir := nilaiAkhir > 76
	lulusAbsen := nilaiAbsen > 75

	lulus := lulusNilaiAkhir && lulusAbsen

	var a, b, c int = 5, 7, 9
	d := a > b && a < c
	e := a > b || a < c
	f := !(a == b && a < c)

	fmt.Println(lulus)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
