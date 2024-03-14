/* If Nestrd Statement
kita bisa memiliki pernyataan if di dalam pernyataan if,
ini disebut dengan nested if.

syntaxnya :
if kondisi1 {
   // kode yang akan dieksekusi jika kondisi1 benar
  if kondisi2 {
     // kode yang akan dieksekusi jika kondisi1 dan kondisi2 benar
  }
}
*/

package main

import "fmt"

func main() {

	angka := 25

	if angka >= 15 {
		fmt.Println("Angka Lebih dari 15")
		if angka > 20 {
			fmt.Println("Angka lebih dari 20")
		}
	} else {
		fmt.Println("Angka Kurang dari 20")
	}
}
