/* Operator Asterisk (*)
Saat kita mengubah variabel pointer, maka yang berubah
hanya variabel tersebut
Semua variabel yang mengacu ke data yang sama tidak akan berubah
jika kita ingin mengubah seluruh variabel yang mengacu pada data tersebut
kita bisa menggunakan operator *

*/

package main

import "fmt"

type Alamat struct {
	Prov, Kab, Kec string
}

func main() {
	alamat1 := Alamat{"Jogja", "Sleman", "Gamping"}
	alamat2 := &alamat1

	alamat2.Kec = "Mlati"

	*alamat2 = Alamat{"Jateng", "Magelang", "Salam"}

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}
