package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {

	csvString := "Misbah,Adrik,Munir\n" +
		"Dafi,Nur,Mahfudhoh\n" +
		"Muhammad,Nur,Adrik"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
