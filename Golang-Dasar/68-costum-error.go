/* Membuat Costum Error
Karena error adalah sebuah interface, jadi ketika kita ingin membuat error sendiri,
kita bisa membuat struct yang mengikuti kontrak dari interface error
*/

package main

import (
	"fmt"
)

// code kontrak error
type validasiError struct {
	Pesan string
}

func (v *validasiError) Error() string {
	return v.Pesan
}

type tidakError struct {
	Pesan string
}

func (v *tidakError) Error() string {
	return v.Pesan
}

// kode menggunakan kostum error
func SaveData(id string, data any) error {
	if id == "" {
		return &validasiError{Pesan: "Validasi Error"}
	}

	if id != "adrik" {
		return &tidakError{Pesan: "Data Tidak Error"}
	}

	return nil
}
func main() {

	err := SaveData("joko", nil)

	if err != nil {
		if validasiError, ok := err.(*validasiError); ok {
			fmt.Println("Validasi Error : ", validasiError.Pesan)
		} else if tidakError, ok := err.(*tidakError); ok {
			fmt.Println("Tidak Ada Error : ", tidakError.Pesan)
		} else {
			fmt.Println("Unknown Error :", err.Error())
		}
	}

}
