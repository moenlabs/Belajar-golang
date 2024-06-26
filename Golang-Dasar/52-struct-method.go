/* Struct Method

struct merupakan type data seperti type data lainnya,
yang mana struct bisa digunakan sebagai parameter untuk function.
namun jika kita ingin menambahkan method ke dalam struct,
sehingga seolah-olah sebuah struct memiliki function

method adalah function yang menempel dalam struct
*/

package main

import "fmt"

type Person struct {
	Nama, Alamat string
	Usia         int8
}

func (orang Person) namaKu(nama string) {

	fmt.Println("Hi", nama, "Namaku adalah :", orang.Nama)
}

func (siswa Person) salam() {
	fmt.Println("Assalamu'alaikum Namaku adalah :", siswa.Nama)
}

func main() {

	adrik := Person{"Adrik", "Jogja", 21}

	fmt.Println(adrik)

	joko := Person{
		Nama:   "Joko",
		Alamat: "Srumbung",
		Usia:   25,
	}

	adrik.namaKu("Adrik")

	joko.salam()

}
