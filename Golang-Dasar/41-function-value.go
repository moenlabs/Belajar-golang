/* function value
Function di golang adalah first class citizen.
artinya function juga merupakan tipe data,
dan juga bisa disimpan di dalam variabel.

menjadikan function sebbagai value dari variabel
tanpa disertai tanda kurung ().

di sini function bisa juga di kirim ke sebuah paramater
di function lain.
*/

package main

import "fmt"

func salam(ucapan string) string {
	return "Assalamu'alaikum " + ucapan
}

func data(nama string) string {
	return "Nama Saya: " + nama
}

func nilai(nilai int) int {
	if nilai <= 90 {
		fmt.Println("Tidak Lulus")
	} else if nilai >= 90 {
		fmt.Println("Anda Lulus")
	}

	return nilai
}

func main() {

	sapaan := salam // variabel sapaan mempunyai nilai function salam
	fmt.Println(sapaan("Ahmad"))

	detail := data
	fmt.Println(detail("Adrik"))

	hasilNilai := nilai
	fmt.Println(hasilNilai(95))

}
