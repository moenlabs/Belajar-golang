/* Nested Loop
kita bisa untuk menempatkan loop (inner loop) di dalam loop lainnya (outer loop).
Di sini, "inner loop" akan dieksekusi satu kali untuk setiap iterasi "outer loop":

*/

package main

import "fmt"

func main() {

	// ini array
	minuman := [2]string{"Kopi", "Teh"}
	makanan := [3]string{"Nasi", "Jagung", "Kacang"}

	for a := 0; a < len(minuman); a++ {
		for b := 1; b < len(makanan); b++ {
			fmt.Println(minuman[a], makanan[b])
		}
	}
}
