/* Saat kita membuat package, kita bisa membuat sebuah function
yang dapat diakses ketika paket dapat diakses

ini sangat cocok ketika contohnya, jika package kita berisi function-function
untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi database

untuk membuat function yang diakses secara otomatis ketika package diakses,
kita cukup membuat function dengan nama init. (sehingga menjadi private)
*/

package main

import (
	"Golang-Dasar/database"
	"fmt"
)

func main() {

	fmt.Println(database.GetDatabase())

}
