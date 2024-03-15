/* Function sebagai parameter
function tidak hanya bisa kita simpan di dalam variable sebagai value
namun function bisa kita gunakan sebagai parameter
untuk function lainnya.
*/

package main

import "fmt"

//

// contoh function parameter 1
func namaMobil(nama string, mobil func(string) string) {
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
