/* Named Return Values

Biasanya saat kita memberi tahu bahwa sebuah function
mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function.
namun, kita juga bisa membuat variable secara langsung di tipe data return functionnya.
*/

package main

import "fmt"

func inventaris() (komputer, laptop, hp int) {
	komputer = 5
	laptop = 2
	hp = 10

	return komputer, laptop, hp
}

func main() {
	a, b, _ := inventaris()
	fmt.Println(a, b)
	fmt.Println(inventaris())
}
