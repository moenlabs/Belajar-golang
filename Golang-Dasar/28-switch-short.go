/* Switch with Short Statemen

sama dengan if, switch mendukung short statemen
sebelum variabel yang akan di cek kondisinya
*/

package main

import "fmt"

func main() {

	switch hari := "jum'at"; hari == "senin" {
	case true:
		fmt.Println("Ini adalah Hari Senin")
	case false:
		fmt.Println("Ini hari Jum'at")
	}
}
