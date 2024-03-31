/* Tipe Data Number
di Golang ada dua tipe data number
	1. integer (bilangan bulat)
	2. folating point (bilangan desimal)

Tipe Data Integer (1) / ada negatifnya
	1. int8 = 	nilai min (-128) nilai max (127)
	2. int16 = nilai min (-32768) nilai max (32767)
	3. int32 = nilai min (-2147483648) nilai max (2147483647)
	4. int64 = nilai min (-9223372036854775808) nilai max (9223372036854775807)
catatan : semakin besar bilangan, semakin besar juga data yang disimpan di memori kita

Tipe Data Integer (2) / tidak ada negatifnya
	1. uint8 = nilai min (0) nilai max (255)
	2. uint16 = nilai min (0) nilai max (65535)
	3. uint32 = nilai min (0) nilai max (4294967295)
	4. uint64 = nilai min (0) nilai max (18446744073709551615)

Tipe Data Floating Point
	1. float32
	2. float64
	3. complex64
	4. complex128

Alias (nama lain)
	1. byte = uint8
	2. rune = int32
	3. int = minimal int32
	4. uint = minimal uint32

*/ 

// Contoh Kode Program

package main

import "fmt"

func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)
	fmt.Println("Bilangan sepuluh", 10)
}
