/* Anonymous Function
sebelumnya setiap  membuat function, kita akan selalu memberikan sebuah nama
pada function tersebut.
namun kadang ada kalanya lebih mudah membuat function secara langsung
di variabel atau parameter tanpa harus membuat function terlebih dahulu
hal ini dinamakan anonymous function atau function tanpa  nama
*/

package main

import "fmt"

// membuat function type

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Anda diblockir", name)
	} else {
		fmt.Println("Selamat Datang", name)
	}
}

func main() {

	// contoh anonymous function 1 dengan
	// dengan membuat anonymous function sebagai nilai dari variabel
	blacklist := func(name string) bool {
		return name == "Asu"
	}

	registerUser("Asu", blacklist)
	// contoh anonymous function 2 dengan
	// langsung mencetak anonymous function
	registerUser("Ahmad", func(name string) bool {
		return name == "Asu"
	})

}
