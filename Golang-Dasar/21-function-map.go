package main

import "fmt"

func main() {

	// ini map
	siswa := map[string]string{"Nama": "Joko", "Alamat": "Mlangi", "Hobi": "Ngaji"}

	fmt.Println(siswa)

	// function len() untuk mendapatkan jumlah data di map
	panjang := len(siswa)

	fmt.Println(panjang)

	// function map[key] untuk mengambil data atau value di map dengan key
	alamat := siswa["Alamat"]

	fmt.Println(alamat)

	//function map[key] = value untuk mengubah data atau value di map dengan key
	siswa["Hobi"] = "Main Game"

	fmt.Println(siswa)

	// function make(map[tipekey]typevalue) untuk membuat map baru
	data := make(map[string]int)
	data["pembelian 1"] = 100
	data["Pembelian 2"] = 200
	data["Pembelian 3"] = 300

	fmt.Println(data)

	//function delete(map, key) untuk menghapus data di map dengan key
	delete(data, "Pembelian 3")

	fmt.Println(data)
}
