/* Blank Identifier
Kadang kadang kita hanya menjalankan init function di package tanpa harus mengeksekusi salah satu
function yang ada di package
Secara default, golang akan komplen ketika package yang di import namun tidak digunakan.
untuk menangani hal tersebut, kita bisa menggunakan blank identifier (_)
sebelum nama package ketika melakukan import
*/

package main

import (
	_ "Golang-Dasar/blank"
	"fmt"
)

func main() {

	fmt.Println("Menggunakan Init Identifier")
}
