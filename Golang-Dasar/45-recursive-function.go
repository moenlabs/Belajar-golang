/* Recursive Function
recursive function adalah function yang memanggil function dirinya sendiri
kadang dalam pekerjaan, kita sering menemui kasus di mana
menggunakan recursive function lebih mudah dibandingkan
tidak menggunakan recursive function.
contoh kasus yang lebih mudah di selesaikan menggunakan recursive
function adalah factorial.
*/

package main

import "fmt"

// function factorial menggunakan loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// function factorial menggunakan recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}
