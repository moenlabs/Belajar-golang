package main

import (
	"fmt"
	"reflect"
)

type Contoh struct {
	Name string `required:"true" max:"10"`
}

type Orang struct {
	Name, Alamat, Email string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Nama", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with value", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(value any) (hasil bool) {
	hasil = true
	a := reflect.TypeOf(value)
	for i := 0; i < a.NumField(); i++ {
		b := a.Field(i)
		if b.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			hasil = data != ""
			if hasil == false {
				return hasil
			}
		}
	}
	return hasil
}

func main() {
	readField(Contoh{"Adrik"})
	readField(Orang{"Muhammad", "", ""})

	orang := Orang{
		Name:   "Muhammad",
		Alamat: "",
		Email:  "moen",
	}

	fmt.Println(IsValid(orang))

}
