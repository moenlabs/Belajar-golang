/* If dengan Short Statement
If mendukung short statemen sebelum kondisi
hal ini sangat cocok untuk membuat statement yang sederhana
sebelum melakukan pengecekan terhadap kondisi

*/

package main

import "fmt"

func main() {

	if hasil := 10; 3*5 == hasil {
		fmt.Println("ups")
	} else if 2*5 == hasil {
		fmt.Println("Great!")
	} else {
		fmt.Println("Belum Beruntung")
	}

}
