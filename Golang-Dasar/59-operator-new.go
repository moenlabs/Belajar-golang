/* Operator New
sebelumnya untuk membuat pointer dengan menggunakan operator &
golang juga memiliki function new yang bisa digunakan untuk membuat pointer
Namun function new hanya mengembalikan pointer ke data kosong
artinya tidak ada data awal
*/

package main

import "fmt"

type Alamat struct {
	Prov, Kab, Kec string
}

func main() {

	rumah1 := new(Alamat)
	rumah2 := rumah1

	rumah2.Kec = "Salam"

	fmt.Println(rumah1)
	fmt.Println(rumah2)
}
