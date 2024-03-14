/* Switch bisa di deklarasikan tanpa
sebuah kondisi
*/

package main

import "fmt"

func main() {
	hari := "jumat"

	switch {
	case hari == "senin":
		fmt.Println("Ini hari Senin")
	case hari == "kamis":
		fmt.Println("Ini hari Kamis")
	default:
		fmt.Println("Ini hari Jum'at")
	}

}
