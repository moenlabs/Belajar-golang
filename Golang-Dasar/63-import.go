/* Import
Secara standar, file di Golang hanya bisa mengakses file golang lainya
yang berbeda dalam package.

jika ingin mengakses file golang yang berbeda di luar package,
maka kita bisa menggunakan Import
*/

package main

import (
	"Golang-Dasar/paket"
	"fmt"
)

func main() {
	transfer := paket.BelajarGolang("Golang")

	fmt.Println(transfer)
}
