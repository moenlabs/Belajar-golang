/* Operasi Matika
Operasi matika digunakan untuk melakukan operasi
matematika secara umum.
+ = penjumlahan
- = pengurangan
* = perkalian
/ = pembagian
% = sisa pembagian (modulo)
++ = menambahkan 1 angka ke variabel
-- = mengurangi 1 angka ke variabel
*/

// Contoh kode
package main

import "fmt"

func main() {

	a := 4
	b := 5
	c := 7
	d := 9

	e := a + b
	f := d - c
	g := a * b * c
	h := g / f

	i := g % h

	j := 10
	j++
	k := 20
	k--

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
