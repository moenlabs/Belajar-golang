package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile(nama string) (string, error) {
	file, err := os.OpenFile(nama, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	var pesan string

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		pesan += string(line) + "\n"
	}

	return pesan, nil
}

func main() {

	hasil, _ := readFile("contoh.log")

	fmt.Println(hasil)

}
