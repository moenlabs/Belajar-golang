package main

import (
	"errors"
	"fmt"
)

var (
	validasiError = errors.New("Validation Error")
	notFound      = errors.New("Not Found")
)

func GetById(id string) error {
	if id == "" {
		return validasiError
	}

	if id != "adrik" {
		return notFound
	}

	// success
	return nil
}

func main() {

	err := GetById("")

	if err != nil {
		if errors.Is(err, validasiError) {
			fmt.Println("validation error")
		} else if errors.Is(err, notFound) {
			fmt.Println("not found error")
		} else {
			fmt.Println("Unknow Error")
		}
	}

}
