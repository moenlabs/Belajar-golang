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
	fmt.Println("Hey", namaDepan, namaBelakang)
}

func main() {
	nama("Ahmad", "Nur")

}
