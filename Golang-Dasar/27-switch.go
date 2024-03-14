/* Switch Statemen
selain if statemen, untuk melakukan percabangan, juga bisa
menggunakan switch statemen
switch statemen sangat sederhada dibandingkan if
biasanya switch statemen digunakan untuk melakukan
pengecekan ke kondisi dalam satu variabel

pernyataan switch digunakan untuk memilih salah satu
dari sekian banyak blok kode yang akan dieksekusi.

switch statemen di Go mirip dengan yang ada di C, C++, Java, JavaScript, dan PHP.
Perbedaannya adalah switch di go hanya menjalankan kasus yang cocok
sehingga tidak memerlukan pernyataan `break`.

Syntax :

switch statemen {
	case x:
   // blok kode
	case y:
   // blok kode
	case z:
	...
	default:
   // blok kode
}
*/

package main

import "fmt"

func main() {

	hari := 6

	switch hari {
	case 1:
		fmt.Println("Ini hari Ahad")
	case 2:
		fmt.Println("Ini hari Senin")
	case 3:
		fmt.Println("Ini hari Selasa")
	case 4:
		fmt.Println("Ini hari Rabu")
	case 5:
		fmt.Println("Ini hari Kamis")
	case 6:
		fmt.Println("Ini hari Jum'at")
	case 7:
		fmt.Println("Ini hari Sabtu")
	default:
		fmt.Println("Anda salah Hari")
	}

}
