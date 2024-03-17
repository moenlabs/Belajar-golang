/* Closure
closure adalah kemampuan sebuah function berinteraksi dengan data-data
di sekitarnya dalam scope yang sama

Penting : Harap gunakan fitur closure ini dengan bijak
		  saat kita membuat aplikasi
*/

package main

import "fmt"

func main() {

	counter := 0
	increment := func() {
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}
