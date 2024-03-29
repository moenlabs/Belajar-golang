/* Parameter Function
saat membuat function, kadang-kadang kita membutuhkan data
dari luar, atau yang disebut parameter.
kita bisa menambahkan parameter di function, bisa lebih dari satu.
parameter tidaklah wajib, jadi kita bisa membuat function tanpa
parameter seperti yang sudah kita buat
namun jika menambahkan parameter di function, maka ketika  memanggil
function tersebut, kita wajib memasukkan data ke parameternya.

sebuah informasi dapat diberikan ke suatu function sebagai parameter.
Parameter bekerja sebagaimana variabel di dalam function.

Parameter dan tipedatanya ditentukan setelah nama function, di dalam tanda kurung ().
kita dapat menambahkan parameter sebanyak yang kita inginkan, cukup pisahkan dengan koma:

Syntax :
func FunctionName(parameter1 type, parameter2 type, parameter3 type) {
  // code yang akan dieksekusi
}

Ingat : membuat function dengan parameter,
harus mengirimkan datanya sesuai dengan tipe datanya
dan jumlah parameternya.
*/

package main

import "fmt"

func nama(namaDepan string, namaBelakang string) {
	fmt.Println("Namaku", namaDepan, namaBelakang)
}

func penjumlahan(a, b int) {
	fmt.Println(a + b)
}

func getMenu(makan, minum string) {

	fmt.Println("makan :", makan, "minum:", minum)
}

func myName(firtsName, lastName string) {
	fmt.Println("Nama Depan: ", firtsName, "\nNama Belakang: ", lastName)
}

func main() {
	nama("Ahmad", "Nur")
	penjumlahan(10, 30)
	getMenu("Pecel Lele", "Jeruk Panas")
	myName("Misbahul", "Munir")
}
