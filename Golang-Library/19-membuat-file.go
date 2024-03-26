package main

import (
	"os"
)

func createNewFile(nama string, pesan string) error {
	file, err := os.OpenFile(nama, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(pesan)

	return nil
}

func main() {

	createNewFile("contoh.log", "ini hanya contoh")

}
