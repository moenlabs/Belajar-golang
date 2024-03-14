/* Multi-case Switch
kita bisa memiliki beberapa nilai untuk setiap case
dalam switch statement.

Syntax :

switch expression {
case x,y:
   // blok kode jika ekspresi dievaluasi ke x atau y
case v,w
   // blok kode jika ekspresi dievaluasi menjadi v atau w
case z
...
default:
   // blok kode jika ekspresi tidak ditemukan dalam case mana pun
}
*/

package main

import "fmt"

func main() {

	hari := 7

	switch hari {
	case 1, 2, 3:
		fmt.Println("Hari Masuk Kerja")
	case 4, 5, 6:
		fmt.Println("Ini hari kerja Terahir")
	case 7:
		fmt.Println("Waktunya Ngoding")
	default:
		fmt.Println("Saatnya Libur Terus")
	}
}
