/* Type assertions
Type assertions merupakan kemampuan merubah tipe data
menjadi tipe data yang diinginkan.
fitur ini sering kali digunakan ketika bertemu
dengan data interface kosong.

Type Assaertion menggunakan Switch

Saat salah menggunakan type assertion, maka bisa
berakibat terjadi panic di aplikasi kita
jika terjadi panic dan tidak ter-recpver,
maka otomatis program kita akan mati.
agar lebih aman, sebaiknya kita menggunakan switch
untuk melakukan type assertion
*/

package main

import "fmt"

func random() any {
	return "Mantab"
}

func main() {

	result := random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	// var resultInt int = result.(int)
	// fmt.Println(resultInt) // panic karena sudah tahu string ko ke int

	// type assertion menggunakan switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unknown", value)
	}

}
