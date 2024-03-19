/* Panic
panic merupakan sebua function yang bisa dipangggil
untuk menghentikan program.
panic function biasanya dipanggil ketika terjadi panic
saat program dijalankan
ketika function panic dipanggul, program akan terhenti,
namun function defer tetap akan di eksekusi.
*/

package main

import "fmt"

func akhir() {
	fmt.Println("Aplikasi sudah berahir")
}

func berjalan(error bool) {
	defer akhir()

	if error {
		panic("terjadi Error")
	} else {
		fmt.Println("Aplikasi Sukses")
	}
}

func main() {
	berjalan(false)

	// berjalan(true)
}
