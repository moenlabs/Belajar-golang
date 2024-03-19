/* ada banyak cara untuk membuat data dari struct
 */

package main

import "fmt"

type Person struct {
	Nama, Alamat string
	Usia         int8
}

func main() {
	// contoh 1
	adrik := Person{
		Nama:   "Adrik",
		Alamat: "Jogja",
		Usia:   20,
	}

	fmt.Println(adrik)

	// contoh 2

	nur := Person{"Nur", "Magelang", 17}
	fmt.Println(nur)

}
