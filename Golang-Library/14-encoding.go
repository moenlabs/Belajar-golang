package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("Muhammad Adrik"))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
	}

	// csv reader
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

	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Adrik", "Muhammad", "Misbah"})
	_ = writer.Write([]string{"Hadafi", "Nur", "Mahfudhoh"})
	_ = writer.Write([]string{"Muhammad", "Misbahul", "Munir"})
	writer.Flush()

}
