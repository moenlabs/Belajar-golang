/* Nil
Biasanya di dalam bahasa pemograman lain, object yang belum diinisialisasi
secara otomatiss nilainya adalah null atau nil.
berbeda dengan golang, saat kita buat variabel dengan tipe data tertentu,
maka secara otomatis akan dibuatkan default valuenya.
Namun, di golang ada itpe data nil, yaitu kosong.
Nil sendiri hanya bisa digunakan di beberapa tipe data,
seperti interface, function, map, slice, pointer dan channel.
*/

package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}

	}
}
func main() {
	data := newMap("Adrik")

	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data["name"])
	}

}
