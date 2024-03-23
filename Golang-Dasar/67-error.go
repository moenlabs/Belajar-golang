/* Error Interface
sebelumnya kita sudah membahas tentang panic, recover dan defer.
nah iktu dilakukan ketika harus menghentikan aplikasinya.
namun ketika kita hanya mau membuat error biasa saja kita bisa menggunakan fitur error.

Golang memiliki interface yang digunakan sebagai kontrak untuk membuat error,
nama interface tersebut adalah error

syntax :

type error interface {
	Error() string
}
simplenya jika ingin membuat error, maka buat dari object interface error ini

Membuat Error
untuk membuat error, kita tidak perlu manual
Golang sudah menyediakan library untuk membuat helper secara mudah
yang dapat di package errors.
*/

package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {

	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
