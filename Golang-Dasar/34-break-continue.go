/* Break & Continue
break dan continue adalah kata kunci yang bisa digunakan dalam perulangan
break : untuk menghentikan seluruh perulangan
continue : untuk menghentikan perulangan yang sedang berjalan,
dan langsung melanjutkan ke perulangan selanjutnya

Pernyataan continue digunakan untuk melewatkan satu atau beberapa iterasi dalam perulangan.
Kemudian dilanjutkan dengan iterasi berikutnya dalam perulangan.
*/

package main

import "fmt"

func main() {

	// contoh break
	for a := 0; a < 7; a++ {
		if a == 4 {
			break
		}

		fmt.Println("Perulangan Ke: ", a)
	}

	// contoh continue
	for b := 1; b < 11; b++ {
		if b%2 == 0 {
			continue
		}

		fmt.Println("Continue ke :", b)
	}

}
