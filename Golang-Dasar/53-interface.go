/* interface adalah tipe data abstrak,
interface tidak memiliki implementasi langsung
sebuah interface berisikan definisi-definisi method
biasanya interface digunakan sebagai kontrak.
arti kontrak adalah bahwa ia harus digunakan oleh struc.

implementasi interface
setiap tipe data yang sesuai dengan kontrak interface,
secara otomatis dianggap sebagai interface tersebut.
sehingga kita tidak perlu mengimplementasikan interface secara manual
hal ini agak berbeda dengan bahasa pemograman lain
yang ketika membuat interface, kita harus menyebutkan
secara eksplisit akan menggunakan interface mana

*/

package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Assalamu'alaikum", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Binatang struct {
	Name string
}

func (binatang Binatang) GetName() string {
	return binatang.Name
}

func main() {

	person := Person{Name: "Andi"}
	SayHello(person)

	animal := Binatang{Name: "Kucing"}
	SayHello(animal)

}
