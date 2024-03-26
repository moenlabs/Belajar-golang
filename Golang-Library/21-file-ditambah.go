package main

import (
	"os"
)

func addTOFile(nama string, pesan string) error {
	file, err := os.OpenFile(nama, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(pesan)

	return nil
}

func main() {

	addTOFile("contoh.log", "\nmenambah file dengan os.O_APPEND ")
}
