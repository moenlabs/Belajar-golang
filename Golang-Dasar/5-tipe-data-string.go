/* Tipe Data String
string adalah tipe data kumpulan karakter
jumlah karakter di dalam string bisa nol sampai tak terhingga
tipe data string di Golang direpresentasikan dengan kata kunci string
nilai data string di Golang selalu diawali dengan karakter "(petik dua) dan
diakhiri  karakter "(petik dua)
*/

// Contoh kode
package  main

import "fmt"

func main() {
	fmt.Println("moen") 
	fmt.Println("Misbah")
	fmt.Println("Munir")

	//function untuk string
	fmt.Println(len("moen")) // hasil 4
	fmt.Println("Misbah"[5]) // hasil 104 (berupa byte)
	fmt.Println("Munir"[0]) // hasil 77 (berupa byte)

}

/* Function untuk String

len("Misbah") = untuk menghitung jumlah karakter misbah
"Misbah"[5] = mengambil karakter misbah nomer 5

*/