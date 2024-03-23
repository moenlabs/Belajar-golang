/* Pointer di Method
Walaupun method akan menempel pada struct,
tapi sebenarnya data struct yang di akses di dalam method adalah pass by value
Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory
karena harus selalu diduplikasi ketika memanggil methodnya.
*/

package main

import "fmt"

type Orang struct {
	Nama string
}

// tanpa pointer sehingga hasilnya tidak berubag
func (orang Orang) Status() {
	orang.Nama = "Sdr. " + orang.Nama
}

// menggunakan pointer

func (data *Orang) Panggilan() {
	data.Nama = "Bapak " + data.Nama
}

func main() {
	adrik := Orang{"Adrik"}

	adrik.Status()

	fmt.Println(adrik.Nama) // hasil cuma Adrik, bukan Sdr.Adrik

	muhammad := Orang{"Muhammad"}

	muhammad.Panggilan()

	fmt.Println(muhammad.Nama)
}
