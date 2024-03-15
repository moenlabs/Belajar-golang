/* Function Type Deklaration
Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
type declaration juga bisa digunakan untuk membuat alias function,
sehingga mempermudah kita menggunakan function sebagai parameter.
*/

package main

import "fmt"

// type function deklaration
type Mobil func(string) string

func namaMobil(nama string, mobil Mobil) {
	fmt.Println("Mobilmu : \n ", mobil(nama))
}

func pilihan(nama string) string {
	if nama == "Brio" {
		return "Mobil Jelek"
	} else if nama == "Pajero" {
		return "Mantab!"
	} else {
		return "Punya Mobil Gak e?"
	}
}

func main() {
	namaMobil("Pajero", pilihan)
	sebutkan := pilihan
	namaMobil("Brio", sebutkan)
	namaMobil("Alya", sebutkan)
}
