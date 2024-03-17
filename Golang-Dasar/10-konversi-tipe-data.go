/* Konversi Tipe Data
di golang kadang butuh melakukan konversi tipe data dari
satu tipe data ke tipe data lain
misal, ingin mengkonversi tipe data int32 ke int63 dan lain-lain.
*/

package main

import (
	"fmt"
)

func main() {

	// merubah int32 ke int64 dan int16
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Misbahul Munir"
	var m = name[0]
	var mString = string(m)

	fmt.Println(name)
	fmt.Println(m)
	fmt.Println(mString)

}
